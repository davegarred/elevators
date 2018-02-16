package main

import (
	"testing"
)

func TestNewElevator(t *testing.T) {
	elevator := NewElevator()
	if elevator.direction != Stopped {
		t.Error()
	}

}
