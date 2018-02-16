package main

import (
	"container/heap"
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


type Riders []*RiderArrival

func NewRiders(r []*RiderArrival) *Riders {
	riders := make(Riders,0)
	riders = append(riders, r...)
	heap.Init(&riders)
	return &riders
}
func (rs *Riders) AddRider(r *RiderArrival) {
	heap.Push(rs, r)
}

func (a Riders) Len() int           { return len(a) }
func (a Riders) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Riders) Less(i, j int) bool { return a[i].arrival < a[j].arrival }

func (rs *Riders) Push(x interface{}) {
	*rs = append(*rs, x.(*RiderArrival))
}

func (rs *Riders) Pop() interface{} {
	pq := *rs
	n := len(pq)
	item := pq[n-1]
	*rs = pq[0 : n-1]
	return item
}

func(rs *Riders) HasRiders() bool {
	return len(*rs) > 0
}