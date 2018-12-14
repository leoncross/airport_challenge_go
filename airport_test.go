package main

import "testing"

func TestAirportLand(t *testing.T) {
	plane := "plane"
	got := Land(plane)
	want := "plane"

	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
