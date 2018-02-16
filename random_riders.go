package main

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GoingUpEvenStrategy(floor int) []*RiderArrival {
	return EvenStrategy(floor, UpRider)
}

func EvenStrategy(floor int, f func(int,int) *RiderArrival) []*RiderArrival {
	arrivals := make([]*RiderArrival,100)
	for i := 0; i<50; i++ {
		baseTime := 60 * i
		arrivals[2*i] = f(floor, baseTime+r.Intn(60))
		arrivals[2*i+1] = f(floor, baseTime+r.Intn(60))
	}
	return arrivals
}

func UpRider(floor int, time int) *RiderArrival {
	return NewRiderArrival(0, floor,time)
}
func DownRider(floor int, time int) *RiderArrival {
	return NewRiderArrival(floor, 0, time)
}
