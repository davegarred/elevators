package main

import (
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

func (a Riders) Len() int           { return len(a) }
func (a Riders) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Riders) Less(i, j int) bool { return a[i].arrival < a[j].arrival }

func (pq *Riders) Push(x interface{}) {
	*pq = append(*pq, x.(*RiderArrival))
}

func (pq *Riders) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

