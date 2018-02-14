package main

type ElevatorCallback struct {
	droppedOffRiders []*RiderArrival
}

func NewElevatorCallback() *ElevatorCallback {
	return &ElevatorCallback{make([]*RiderArrival,0)}
}

func (c *ElevatorCallback) DropOffRider(rider *RiderArrival) {
	c.droppedOffRiders = append(c.droppedOffRiders, rider)
}

func (c *ElevatorCallback) GetDroppedOffRiders() []*RiderArrival {
	return c.droppedOffRiders
}
