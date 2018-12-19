package weather

import (
	"testing"
	"time"
)

var stubbed_sunny = func() time.Time { return time.Unix(0, 0) }
var stubbed_stormy = func() time.Time { return time.Unix(5, 0) }

func TestAirportWeather_S(t *testing.T) {
	actual := WeatherCheck(stubbed_sunny)
	want := "sunny"

	if want != actual {
		t.Errorf("actual '%s' want '%s'", actual, want)
	}
}

func TestAirportWeather_stormy(t *testing.T) {
	actual := WeatherCheck(stubbed_stormy)
	want := "stormy"

	if want != actual {
		t.Errorf("actual '%s' want '%s'", actual, want)
	}
}
