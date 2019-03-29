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

var lc *localClock = new(localClock)

func LocalClock() Clock {
	return lc
}

type Uptimer interface {
	Uptime() float64
}

type mockUptimer struct {
	startAt time.Time
}

func newMockUptimer() *mockUptimer {
	t := &mockUptimer{}
	t.startAt = time.Now()
	return t
}

func (t *mockUptimer) Uptime() float64 {
	return time.Since(t.startAt).Seconds()
}

var mt = newMockUptimer()

func MockUptimer() Uptimer {
	return mt
}
