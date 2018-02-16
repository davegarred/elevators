package main

import (
	"testing"
	"fmt"
	"container/heap"
)

func TestRider(t *testing.T) {
	arrivals := make([]*RiderArrival,0)
	arrivals = append(arrivals,NewRiderArrival(0, 5, 13))
	arrivals = append(arrivals,NewRiderArrival(0, 2, 31))
	arrivals = append(arrivals,NewRiderArrival(0, 8, 19))
	arrivals = append(arrivals,NewRiderArrival(0, 1, 4))

	r := NewRiders(arrivals)
	r.AddRider(NewRiderArrival(0, 4, 15))

	validateSorted(t,r)
}

func TestFloorList(t *testing.T) {
	arrivals := GoingUpEvenStrategy(2)
	arrivals = append(arrivals, GoingUpEvenStrategy(3)...)
	arrivals = append(arrivals, GoingUpEvenStrategy(4)...)
	r := NewRiders(arrivals)
	validateSorted(t,r)
}

func validateSorted(t *testing.T, rs *Riders) {
	time := 0
	for rs.HasRiders() {
		r := heap.Pop(rs).(*RiderArrival)
		if time > r.arrival {
			t.Errorf("Riders not correctly sorted by date, arrival at %d should not be before last at %d", r.arrival, time)
		}
		//r.print()
		time = r.arrival
	}
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
