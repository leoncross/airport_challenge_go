package main

import (
	"airport/plane"
)

func Land(plane *plane.Plane) string {
	plane.Status = "landed"
	return plane.Status
}

func Takeoff(plane *plane.Plane) string {
	plane.Status = "flying"
	return plane.Status
}

func main() {}

// As an air traffic controller
// To ensure safety
// I want to prevent takeoff when weather is stormy
