package main

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GoingUpEvenStrategy(floor int) Riders {
	return EvenStrategy(floor, UpRider)
}

func EvenStrategy(floor int, f func(int,int) *RiderArrival) Riders {
	riders := make(Riders,100)
	for i := 0; i<50; i++ {
		baseTime := 60 * i
		riders[2*i] = f(floor, baseTime+r.Intn(60))
		riders[2*i+1] = f(floor, baseTime+r.Intn(60))
	}
	return riders
}

func UpRider(floor int, time int) *RiderArrival {
	return NewRiderArrival(0, floor,time)
}
func DownRider(floor int, time int) *RiderArrival {
	return NewRiderArrival(floor, 0, time)
}
