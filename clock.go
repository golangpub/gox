package gox

import "time"

type Clock interface {
	Now() time.Time
}

type localClock struct {
}

func (c *localClock) Now() time.Time {
	return time.Now()
}

func LocalClock() Clock {
	return new(localClock)
}

type Uptimer interface {
	Uptime() int64
}
