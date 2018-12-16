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
