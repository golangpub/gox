package types

import "sync/atomic"

type Counter struct {
	count uint32
}

func (c *Counter) Next() uint32 {
	atomic.AddUint32(&c.count, 1)
	return c.count
}

var DefaultCounter = &Counter{}
