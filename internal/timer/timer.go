package timer

import "time"

func Timed(work func() error) (time.Duration, error) {
	start := time.Now()
	err := work()
	return time.Since(start), err
}
