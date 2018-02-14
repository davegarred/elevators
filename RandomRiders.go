package main

import (
	"math/rand"
)

func GoingUpEvenStrategy(floor int) Riders {
	riders := make(Riders,100)
	for i := 0; i<50; i++ {
		baseTime := 60 * i
		riders[2*i] = NewRiderArrival(0,floor, baseTime + rand.Intn(60))
		riders[2*i+1] = NewRiderArrival(0,floor, baseTime + rand.Intn(60))
	}
	return riders
}
