package utilities

import "time"

//go:generate mockgen -source=Clock.go -destination=./mocks/interfaces.go -package=mocks

type Clock interface {
	Now() time.Time
}

type NormalClock struct{}

func (c NormalClock) Now() time.Time {
	return time.Now()
}
