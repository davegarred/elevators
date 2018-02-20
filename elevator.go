package main

import ()

const (
	TransportingUp   = Status(iota)
	TransportingDown = Status(iota)
	Stopped          = Status(iota)
	Repositioning    = Status(iota)

	defaultCapacity = 10

	defaultInitialFloorSpeed = 2 * seconds
	defaultAddtlFloorSpeed   = 1 * seconds
	defaultDoorOpenCloseTime = 6 * seconds
)

type ElevatorType struct {
	capacity             int
	initialFloorSpeed    IntTime
	additionalFloorSpeed IntTime
	doorOpenTime         IntTime
}

func NewDefaultElevator() *ElevatorType {
	return &ElevatorType{
		capacity:             defaultCapacity,
		initialFloorSpeed:    defaultInitialFloorSpeed,
		additionalFloorSpeed: defaultAddtlFloorSpeed,
		doorOpenTime:         defaultDoorOpenCloseTime,
	}
}

type IntTime int
type Floor int
type Status int

type Elevator struct {
	*ElevatorType
	callback       *ElevatorCallback
	necessaryStops []Floor
	passengers     []*Rider
	direction      Status
	currentFloor   Floor
	targetFloor    Floor
	status         Status
	arrivalTime    IntTime
	busyUntil      IntTime
}

func NewElevator(callback *ElevatorCallback) *Elevator {
	elevatorType := NewDefaultElevator()
	return &Elevator{
		ElevatorType:   elevatorType,
		callback:       callback,
		necessaryStops: make([]Floor, 0, elevatorType.capacity),
		passengers:     make([]*Rider, 0, elevatorType.capacity),
		direction:      Stopped,
		arrivalTime:    0,
		currentFloor:   0,
		targetFloor:    0,
		status:         Stopped,
	}
}

func (e *Elevator) SetTargetFloor(time IntTime, floor Floor) IntTime {
	e.status = Repositioning
	return e.travelToFloor(floor, time)
}

func (e *Elevator) travelToFloor(floor Floor, time IntTime) IntTime {
	floorDelta := e.currentFloor - floor
	if floorDelta == 0 {
		return 0
	} else if floorDelta < 0 {
		floorDelta = -floorDelta
	}
	e.targetFloor = floor
	travelTime := e.ElevatorType.initialFloorSpeed + IntTime(floorDelta-1)*e.ElevatorType.additionalFloorSpeed
	e.busyUntil = time + travelTime
	return e.busyUntil
}

func (e *Elevator) Update(time IntTime) (IntTime, Status) {
	if time < e.busyUntil {
		return e.busyUntil, e.status
	}
	e.currentFloor = e.targetFloor
	e.dropOffRiders(e.targetFloor)
	nextFloor := e.findNextFloor()
	if nextFloor != e.currentFloor {
		e.busyUntil = e.ElevatorType.doorOpenTime + e.travelToFloor(nextFloor, time)
	} else {
		e.busyUntil = time
		e.status = Stopped
	}
	return e.busyUntil, e.status
}

func (e *Elevator) dropOffRiders(floor Floor) {
	currentPassengers := e.passengers
	e.passengers = make([]*Rider, 0, e.ElevatorType.capacity)
	for _, rider := range currentPassengers {
		if rider.destination == floor {
			e.callback.DropOffRider(rider)
		} else {
			e.passengers = append(e.passengers, rider)
		}
	}
}

func (e *Elevator) findNextFloor() Floor {
	nextFloor := e.currentFloor
	if e.status == TransportingUp {
		for _, rider := range e.passengers {
			d := rider.destination
			if d > e.currentFloor && (d < nextFloor || e.currentFloor == nextFloor) {
				nextFloor = d
			}
		}
		return nextFloor
	} else if e.status == TransportingDown {
		for _, rider := range e.passengers {
			d := rider.destination
			if d < e.currentFloor && (d > nextFloor || e.currentFloor == nextFloor) {
				nextFloor = d
			}
		}
		return nextFloor
	}
	if len(e.passengers) > 0 {
		panic("Elevator has passengers but can not determine the next floor if status is unknown")
	}
	return e.currentFloor
}

func (e *Elevator) Pickup(time IntTime, riders Riders, status Status) (IntTime, Status) {
	if len(riders) == 0 {
		return time, e.status
	}
	e.status = status
	currentNumPass := len(e.passengers)
	for _, rider := range riders {
		if currentNumPass == e.ElevatorType.capacity {
			break
		}
		e.passengers = append(e.passengers, rider)
		currentNumPass++
	}
	nextFloor := e.findNextFloor()
	e.busyUntil = e.ElevatorType.doorOpenTime + e.travelToFloor(nextFloor, time)
	return e.busyUntil, e.status
}

// Match w/elevator = e.idle || e.canReach
// canReach = (e.stopped || e.canStop(floor)) && (e.nextStop == null || e.nextStop after floor)
