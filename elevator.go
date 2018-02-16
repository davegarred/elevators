package main

import ()

const (
	TransportingUp   = iota
	TransportingDown = iota
	Stopped          = iota
	Repositioning    = iota

	defaultCapacity          = 10
	defaultInitialFloorSpeed = 2000
	defaultAddtlFloorSpeed   = 1000
	defaultDoorOpenTime      = 6000
)

type ElevatorType struct {
	capacity             int
	initialFloorSpeed    int
	additionalFloorSpeed int
	doorOpenTime         int
}

func NewDefaultElevator() *ElevatorType {
	return &ElevatorType{
		capacity:             defaultCapacity,
		initialFloorSpeed:    defaultInitialFloorSpeed,
		additionalFloorSpeed: defaultAddtlFloorSpeed,
		doorOpenTime:         defaultDoorOpenTime,
	}
}

type Status int

type Elevator struct {
	*ElevatorType
	necessaryStops  []int
	direction       int
	currentLoad     int
	currentFloor    int
	targetFloor     int
	status Status
	arrivalTime     int
	lastCommandTime int
	busyUntil       int
}

func NewElevator() *Elevator {
	elevatorType := NewDefaultElevator()
	return &Elevator{
		ElevatorType:    elevatorType,
		necessaryStops:  make([]int, elevatorType.capacity),
		direction:       Stopped,
		arrivalTime:     0,
		currentFloor:    0,
		targetFloor:     0,
		status: Stopped,
	}
}

func (e *Elevator) Direction() int {
	return e.direction
}

func (e *Elevator) movingUp() bool {
	switch {
	case e.targetFloor > e.currentFloor:
		return true
	case e.targetFloor < e.currentFloor:
		return false
	default:
		panic("Programming error: Attempt to determine direction of non-moving elevator")
	}
}
func (e *Elevator) position() int {
	if e.currentFloor != e.targetFloor {
		panic("Programming error: Attempt to check position of a moving elevator")
	}
	return e.currentFloor
}

func (e *Elevator) SetTargetFloor(time int, floor int) int {
	floorDelta := e.currentFloor - floor
	if floorDelta == 0 {
		return 0
	} else if floorDelta < 0 {
		floorDelta = -floorDelta
	}

	travelToFloor(e, floor, floorDelta, time)
	return e.busyUntil
}
func travelToFloor(e *Elevator, floor int, floorDelta int, time int) {
	e.targetFloor = floor
	travelTime := e.ElevatorType.initialFloorSpeed + (floorDelta-1)*e.ElevatorType.additionalFloorSpeed
	e.lastCommandTime = time
	e.busyUntil = time + travelTime
	e.status = Repositioning
}

func (e *Elevator) Update(time int) (int,Status) {
	if time < e.busyUntil {
		return e.busyUntil, e.status
	}

	e.lastCommandTime = time
	e.status = Stopped
	return e.lastCommandTime, e.status
}
//func (e *Elevator) Pickup(time int, []*RiderArrival, Status) int {
//	return
//}

// Match w/elevator = e.idle || e.canReach
// canReach = (e.stopped || e.canStop(floor)) && (e.nextStop == null || e.nextStop after floor)
