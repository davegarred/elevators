package main

import (
	"fmt"
)

type RiderArrival struct {
	origin      int
	destination int
	arrival     int
}

func NewRiderArrival(o int, d int, a int) *RiderArrival {
	return &RiderArrival{o, d, a}
}

func (a *RiderArrival) goingUp() bool {
	return a.destination > a.origin
}

func (a *RiderArrival) print() {
	var direction string
	if a.goingUp() {
		direction = "up"
	} else {
		direction = "down"
	}
	fmt.Printf("Going %s from floor %d to floor %d, arrived at %d\n", direction, a.origin, a.destination, a.arrival)
}

type Riders []*RiderArrival

func (a Riders) Len() int           { return len(a) }
func (a Riders) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Riders) Less(i, j int) bool { return a[i].arrival < a[j].arrival }

func (rs Riders) print() {
	for _,r := range rs {
		r.print()
	}
}
