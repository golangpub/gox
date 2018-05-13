package types

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

// Change to int64, as https://github.com/golang/go/issues/12401 is fixed in golang v1.6
type ID int64

// JSON中数字表示为double，double整数部分最大值为2^53，由于部分JSON库默认不支持int64，因此控制在53bit内比较好
// id由time+shard+seq组成
// 若业务多可扩充shard，并发高可扩充seq. 由于time在最高位,故扩展后的id集合与原id集合不会出现交集,可保持全局唯一

const ShardBitSize = 2 // 最多4个shard

const FastSeqBitSize = 4    //每个shard每ms不能超过16次调用
const DefaultSeqBitSize = 2 //每个shard每ms不能超过4次调用
const SlowSeqBitSize = 3    //每个shard每s不能超过8次调用

var epoch time.Time
var MonoFastIDGenerator *IDGenerator
var MonoDefaultIDGenerator *IDGenerator
var MonoSlowIDGenerator *IDGenerator

func init() {
	epoch = time.Date(2017, time.January, 2, 15, 4, 5, 0, time.UTC)
	MonoFastIDGenerator = NewIDGenerator(0, 0, FastSeqBitSize, false)
	MonoDefaultIDGenerator = NewIDGenerator(0, 0, DefaultSeqBitSize, false)
	MonoSlowIDGenerator = NewIDGenerator(0, 0, SlowSeqBitSize, true)
}

// NewID returns new ID created by default id generator
func NewID() ID {
	return MonoDefaultIDGenerator.NewID()
}

func NewFastID() ID {
	return MonoFastIDGenerator.NewID()
}

func NewSlowID() ID {
	return MonoSlowIDGenerator.NewID()
}

// ShortString returns a short representation of id
func (i ID) ShortString() string {
	if i < 0 {
		panic("invalid id")
	}
	var bytes [16]byte
	k := int64(i)
	n := 15
	for {
		j := k % 62
		switch {
		case j <= 9:
			bytes[n] = byte('0' + j)
		case j <= 35:
			bytes[n] = byte('A' + j - 10)
		default:
			bytes[n] = byte('a' + j - 36)
		}
		k /= 62
		if k == 0 {
			return string(bytes[n:])
		}
		n--
	}
}

func ParseShortID(s string) (ID, error) {
	if len(s) == 0 {
		return 0, errors.New("parse error")
	}

	var bytes = []byte(s)
	var k int64
	var v int64
	for _, b := range bytes {
		switch {
		case b >= '0' && b <= '9':
			v = int64(b - '0')
		case b >= 'A' && b <= 'Z':
			v = int64(10 + b - 'A')
		case b >= 'a' && b <= 'z':
			v = int64(36 + b - 'a')
		default:
			return 0, errors.New("parse error")
		}
		k = k*62 + v
	}
	return ID(k), nil
}

type IDGenerator struct {
	seq          ID
	shardID      ID
	seqBitSize   uint
	shardBitSize uint
	useSecond    bool
}

func NewIDGenerator(shardID, shardBitSize, seqBitSize uint, useSecond bool) *IDGenerator {
	if seqBitSize < 1 || seqBitSize > 16 {
		panic("seqBitSize should be [1,16]")
	}

	if shardBitSize < 0 || shardBitSize > 8 {
		panic("shardBitSize should be [0,8]")
	}

	if shardBitSize+seqBitSize >= 20 {
		panic("shardBitSize + seqBitSize should be less than 20")
	}

	if shardBitSize > 0 {
		if shardID < 0 || shardID > (1<<shardBitSize)-1 {
			panic(fmt.Sprint("shardID must be [ 0,", (1<<shardBitSize)-1, "]"))
		}
	} else {
		//		log.Info("shardBitSize is 0, skip shardID")
	}

	g := &IDGenerator{}
	g.seq = 1
	if g.shardBitSize > 0 {
		g.shardID = ID(shardID << seqBitSize)
	}
	g.seqBitSize = seqBitSize
	g.shardBitSize = shardBitSize
	g.useSecond = useSecond
	return g
}

func (g *IDGenerator) Clone() *IDGenerator {
	return &*g
}

func (g *IDGenerator) NewID() ID {
	seq := ID(atomic.AddInt64((*int64)(&g.seq), 1))
	id := seq % ID(1<<g.seqBitSize)
	id |= g.shardID

	d := time.Since(epoch)
	var timestamp int64
	if g.useSecond {
		timestamp = d.Nanoseconds() / 1e9 //seconds
	} else {
		timestamp = d.Nanoseconds() / 1e6 //milliseconds
	}
	id |= ID(timestamp << (g.seqBitSize + g.shardBitSize))
	return ID(id)
}
