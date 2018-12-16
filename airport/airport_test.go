package main

import (
	"airport/plane"
	"testing"
)

func TestAirportLand(t *testing.T) {
	var p1 = plane.Plane{"flying"}

	actual := Land(&p1)
	want := "landed"

	if want != actual {
		t.Errorf("actual '%s' want '%s'", actual, want)
	}
}

func TestAirportTakeoff(t *testing.T) {
	var p1 = plane.Plane{"landed"}

	actual := Takeoff(&p1)
	want := "flying"

	if want != actual {
		t.Errorf("actual '%s' want '%s'", actual, want)
	}
}
