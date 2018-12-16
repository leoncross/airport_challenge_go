package weather

import (
	"testing"
	"time"
)

var stubbed = func() time.Time { return time.Unix(0, 0) }

func TestAirportWeather(t *testing.T) {
	actual := WeatherCheck(stubbed)
	want := "sunny"

	if want != actual {
		t.Errorf("actual '%s' want '%s'", actual, want)
	}
}
