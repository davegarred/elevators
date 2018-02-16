package main

import (
	"testing"
)

func TestNewElevator(t *testing.T) {
	time := 0
	elevator := NewElevator()
	if elevator.direction != Stopped {
		t.Error()
	}
	next := elevator.SetTargetFloor(time,3)
	if next != 4000 {
		t.Error(next)
	}
	if next,status := elevator.Update(3000); next != 4000 || status != Repositioning {
		t.Errorf("Expected 4000-%d, found %d-%d", Repositioning, next,status)
	}
	if next,status := elevator.Update(4000); next != 4000 || status != Stopped  {
		t.Errorf("Expected 4000-%d, found %d-%d", Stopped, next,status)
	}
	if next,status := elevator.Update(4001); next != 4001 || status != Stopped  {
		t.Errorf("Expected 4001-%d, found %d-%d", Stopped, next, status)
	}

}
