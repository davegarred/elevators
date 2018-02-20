package main

type ElevatorCallback struct {
	droppedOffRiders []*Rider
}

func NewElevatorCallback() *ElevatorCallback {
	return &ElevatorCallback{make([]*Rider,0)}
}

func (c *ElevatorCallback) DropOffRider(rider *Rider) {
	c.droppedOffRiders = append(c.droppedOffRiders, rider)
}

func (c *ElevatorCallback) GetDroppedOffRiders() []*Rider {
	return c.droppedOffRiders
}
