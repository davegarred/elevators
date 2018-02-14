package main

import (
	"testing"
)

func TestMoveElevator(t *testing.T) {
	callback := NewElevatorCallback()
	elevator := NewElevator(callback)
	if elevator.direction != Stopped {
		t.Error()
	}
	next := elevator.SetTargetFloor(0,3)
	if next != 4 * seconds {
		t.Error(next)
	}

	testUpdateAndCheckTime(elevator, 3 * seconds, 4 * seconds, Repositioning, t)
	testUpdateAndCheckTime(elevator, 4 * seconds, 4 * seconds, Stopped, t)
	testUpdateAndCheckTime(elevator, 4001 * milliseconds, 4001 * milliseconds, Stopped, t)
}

func TestPickupOne(t *testing.T) {
	callback := NewElevatorCallback()
	elevator := NewElevator(callback)
	riders := make([]*RiderArrival, 1)
	riders[0] = NewRiderArrival(0,2, 0)

	if next,status := elevator.Pickup(0, riders, TransportingUp); next != (6 + 3) * seconds || status != TransportingUp {
		t.Fatalf("Expected: %v-%v, found: %v-%v\n", 9*seconds, TransportingUp, next, status)
	}
	assertCorrectNumberOfRidersDroppedOff(callback, 0, t)

	testUpdateAndCheckTime(elevator, 8*seconds, 9*seconds, TransportingUp, t)
	assertCorrectNumberOfRidersDroppedOff(callback, 0, t)

	testUpdateAndCheckTime(elevator, 9*seconds, 9*seconds, Stopped, t)
	assertCorrectNumberOfRidersDroppedOff(callback, 1, t)

}

func TestPickupTwo_differentFloors(t *testing.T) {
	callback := NewElevatorCallback()
	elevator := NewElevator(callback)
	riders := make([]*RiderArrival, 2)
	riders[0] = NewRiderArrival(0,2, 0)
	riders[1] = NewRiderArrival(0,4, 0)

	elevator.Pickup(0, riders, TransportingUp)

	testUpdateAndCheckTime(elevator, 8*seconds, 9*seconds, TransportingUp, t)

	testUpdateAndCheckTime(elevator, 9*seconds, 18*seconds, TransportingUp, t)
	assertCorrectNumberOfRidersDroppedOff(callback, 1, t)

	testUpdateAndCheckTime(elevator, 17*seconds, 18*seconds, TransportingUp, t)

	testUpdateAndCheckTime(elevator, 18*seconds, 18*seconds, Stopped, t)
	assertCorrectNumberOfRidersDroppedOff(callback, 2, t)

}

func testUpdateAndCheckTime(elevator *Elevator, testTime IntTime, expectedTime IntTime, expectedStatus Status, t *testing.T) {
	if next, status := elevator.Update(testTime); next != expectedTime || status != expectedStatus {
		t.Fatalf("Expected %d-%d, found %d-%d", expectedTime, expectedStatus, next, status)
	}
}

func assertCorrectNumberOfRidersDroppedOff(callback *ElevatorCallback, expectedRiders int, t *testing.T) {
	if len(callback.GetDroppedOffRiders()) > expectedRiders {
		t.Fatalf("Expected %d riders dropped off yet, found %d\n", expectedRiders, len(callback.GetDroppedOffRiders()))
	}
}
