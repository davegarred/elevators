package main

import (
)
const (
	Up = iota
	Down = iota
	Stopped = iota

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

type Elevator struct {
	*ElevatorType
	necessaryStops []int
	direction int
	currentLoad  int
	currentFloor int
	targetFloor  int
	arrivalTime  int
}
func NewElevator() *Elevator {
	elevatorType := NewDefaultElevator()
	return &Elevator{
		ElevatorType:elevatorType,
		necessaryStops: make([]int,elevatorType.capacity),
		direction:Stopped,
		arrivalTime:0,
		currentFloor:0,
		targetFloor:0,
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

func (e *Elevator) SetTargetFloor(floor int) int {
	floorDelta := e.currentFloor - floor
	if floorDelta == 0 {
		return 0
	} else if floorDelta < 0 {
		floorDelta = -floorDelta
	}

	e.targetFloor = floor
	time := e.ElevatorType.initialFloorSpeed + (floorDelta - 1) * e.ElevatorType.additionalFloorSpeed
	return time
}

// Match w/elevator = e.idle || e.canReach
// canReach = (e.stopped || e.canStop(floor)) && (e.nextStop == null || e.nextStop after floor)