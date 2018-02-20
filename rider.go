package main

import (
	"container/heap"
)

type Rider struct {
	origin      Floor
	destination Floor
	arrival     IntTime
}

func NewRiderArrival(origin Floor, destination Floor, arrivalTime IntTime) *Rider {
	return &Rider{origin, destination, arrivalTime}
}


type Riders []*Rider

func NewRiders(r []*Rider) *Riders {
	riders := make(Riders,0)
	riders = append(riders, r...)
	if len(riders) > 0 {
		heap.Init(&riders)
	}
	return &riders
}
func (rs *Riders) AddRider(r *Rider) {
	heap.Push(rs, r)
}

func (a Riders) Len() int           { return len(a) }
func (a Riders) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Riders) Less(i, j int) bool { return a[i].arrival < a[j].arrival }

func (rs *Riders) Push(x interface{}) {
	*rs = append(*rs, x.(*Rider))
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