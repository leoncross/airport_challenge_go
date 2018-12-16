package weather

import (
	"math/rand"
	"time"
)

func WeatherCheck(now func() time.Time) string {
	if now == nil {
		now = time.Now
	}
	rand.Seed(now().Unix())
	reasons := []string{
		"stormy",
		"sunny",
		"sunny",
	}

	n := rand.Int() % len(reasons)
	return reasons[n]
}
