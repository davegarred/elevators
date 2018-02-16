package main

import (
	"fmt"
	"testing"
)

func TestNewElevator(t *testing.T) {
	elevator := NewElevator()
	fmt.Printf("%+v\n", elevator)

}
