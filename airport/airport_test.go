// https://github.com/stretchr/testify

package main

import (
	"airport/plane"
	"testing"
)

func TestAirportLand(t *testing.T) {
	t.Run("will try to land a planed", func(t *testing.T) {
		var p1 = plane.Plane{"flying"}
		actual := Land(&p1)
		want := "landed"

		if want != actual {
			t.Errorf("actual '%s' want '%s'", actual, want)
		}
	})
}
func TestAirportLandStormy(t *testing.T) {
	t.Run("will try to land a planed in stormy weather", func(t *testing.T) {
		var p1 = plane.Plane{"flying"}
		actual := Land(&p1)
		want := "cannot land in stormy weather"

		if want != actual {
			t.Errorf("actual '%s' want '%s'", actual, want)
		}
	})
}
func TestAirportTakeoff(t *testing.T) {
	t.Run("will try to take off a plane", func(t *testing.T) {
		var p1 = plane.Plane{"landed"}

		actual := Takeoff(&p1)
		want := "flying"

		if want != actual {
			t.Errorf("actual '%s' want '%s'", actual, want)
		}
	})
	t.Run("will try to takeoff a planed in stormy weather", func(t *testing.T) {
		var p1 = plane.Plane{"landed"}

		actual := Takeoff(&p1)
		want := "cannot takeoff in stormy weather"

		if want != actual {
			t.Errorf("actual '%s' want '%s'", actual, want)
		}
	})
}
