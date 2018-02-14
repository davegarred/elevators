package main

import (
	"math"
)
const (
	defaultCapacity          = 10
	defaultInitialFloorSpeed = 2
	defaultAddtlFloorSpeed   = 1
	defaultDoorOpenTime      = 6
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
	currentLoad  int
	currentFloor int
	targetFloor  int
	arrivalTime  int
}
func NewElevator() *Elevator {
	elevatorType := NewDefaultElevator()
	return &Elevator{
		ElevatorType:elevatorType,
		arrivalTime:0,
		currentFloor:0,
		targetFloor:0,
	}
}

func (e *Elevator) moving() bool {
	return e.targetFloor != e.currentFloor
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
	floorDelta := math.Abs(e.currentFloor - floor)
	if floorDelta == 0 {
		return 0
	} else if floorDelta < 0 {
		floorDelta = -floorDelta
	}

	e.targetFloor = floor
	time := e.ElevatorType.initialFloorSpeed + (floorDelta - 1) * e.ElevatorType.additionalFloorSpeed
	return time
}

