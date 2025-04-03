package mocking

import "time"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
