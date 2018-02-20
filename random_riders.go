package main

import (
	"math/rand"
	"time"
)

const (
	seconds = IntTime(1000)
	milliseconds = IntTime(1)
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GoingUpEvenStrategy(floor Floor) []*Rider {
	return EvenStrategy(floor, UpRider)
}

func EvenStrategy(floor Floor, f func(Floor, IntTime) *Rider) []*Rider {
	arrivals := make([]*Rider, 100)
	for i := 0; i < 50; i++ {
		baseTime := IntTime(60 * i) * seconds
		arrivals[2*i] = f(floor, baseTime+IntTime(r.Intn(60))*seconds)
		arrivals[2*i+1] = f(floor, baseTime+IntTime(r.Intn(60))*seconds)
	}
	return arrivals
}

func UpRider(floor Floor, time IntTime) *Rider {
	return NewRiderArrival(0, floor, time)
}
func DownRider(floor Floor, time IntTime) *Rider {
	return NewRiderArrival(floor, 0, time)
}
