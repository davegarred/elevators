package main

import (
	"testing"
	"sort"
)

func TestRider(t *testing.T) {
	riders := make(Riders,0)
	riders = append(riders,NewRiderArrival(0, 5, 13))
	riders = append(riders,NewRiderArrival(0, 2, 31))
	riders = append(riders,NewRiderArrival(0, 8, 19))
	riders = append(riders,NewRiderArrival(0, 1, 4))
	riders.print()

	sort.Sort(riders)
	riders.print()
}

func TestFloorList(t *testing.T) {
	riders := GoingUpEvenStrategy(2)
	riders = append(riders, GoingUpEvenStrategy(3)...)
	riders = append(riders, GoingUpEvenStrategy(4)...)
	sort.Sort(riders)
	riders.print()
}
