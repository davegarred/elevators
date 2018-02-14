package main

import (
	"testing"
	"sort"
	"fmt"
)

func TestRider(t *testing.T) {
	riders := make(Riders,0)
	riders = append(riders,NewRiderArrival(0, 5, 13))
	riders = append(riders,NewRiderArrival(0, 2, 31))
	riders = append(riders,NewRiderArrival(0, 8, 19))
	riders = append(riders,NewRiderArrival(0, 1, 4))
	riders.print()

	fmt.Println()
	fmt.Println("Sorting...")
	sort.Sort(riders)
	riders.print()
	validateSorted(t,riders)
}

func TestFloorList(t *testing.T) {
	riders := GoingUpEvenStrategy(2)
	riders = append(riders, GoingUpEvenStrategy(3)...)
	riders = append(riders, GoingUpEvenStrategy(4)...)
	sort.Sort(riders)
	validateSorted(t,riders)
	 //riders.print()
}

func validateSorted(t *testing.T, rs Riders) {
	time := 0
	for _,r := range rs {
		if time > r.arrival {
			t.Errorf("Riders not correctly sorted by date, arrival at %d should not be before last at %d", r.arrival, time)
		}
		time = r.arrival
	}
}
