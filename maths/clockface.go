package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	/* 1 second is 2pi/60 radians
	-> float64(t.Second()) * (math.Pi / 30)
	-> re-arrange to avoid first dividing down and then multiplying up,
		as this would results in a very inaccurate result because of
		inaccurate floating point arithmetic
	*/
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
