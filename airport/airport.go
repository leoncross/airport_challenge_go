package main

import (
	"airport/plane"
	"airport/weather"
	"fmt"
)

func Land(plane *plane.Plane) string {
	if Weather() == false {
		return ("cannot land in stormy weather")
	}
	plane.Status = "landed"
	return plane.Status
}

func Takeoff(plane *plane.Plane) string {
	if Weather() == false {
		return ("cannot takeoff in stormy weather")
	}
	plane.Status = "flying"
	return plane.Status
}

func Weather() bool {
	if weather.WeatherCheck(nil) == "sunny" {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(Weather())
}

// As an air traffic controller
// To ensure safety
// I want to prevent takeoff when weather is stormy
