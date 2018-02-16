package main

import (
	"testing"
	"sort"
	"fmt"
	"container/heap"
)

func TestRider(t *testing.T) {
	riders := make(Riders,0)
	riders = append(riders,NewRiderArrival(0, 5, 13))
	riders = append(riders,NewRiderArrival(0, 2, 31))
	riders = append(riders,NewRiderArrival(0, 8, 19))
	riders = append(riders,NewRiderArrival(0, 1, 4))

	heap.Init(&riders)
	heap.Push(&riders, NewRiderArrival(0, 4, 15))
	validateSorted(t,riders)
}

func TestFloorList(t *testing.T) {
	riders := GoingUpEvenStrategy(2)
	riders = append(riders, GoingUpEvenStrategy(3)...)
	riders = append(riders, GoingUpEvenStrategy(4)...)
	sort.Sort(riders)
	validateSorted(t,riders)
}

func validateSorted(t *testing.T, rs Riders) {
	time := 0
	for len(rs) > 0 {
		r := heap.Pop(&rs).(*RiderArrival)
		if time > r.arrival {
			t.Errorf("Riders not correctly sorted by date, arrival at %d should not be before last at %d", r.arrival, time)
		}
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
