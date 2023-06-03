package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 12
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
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / (float64(t.Minute()))))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	// get the remainder of the hour divided by 12, as this is a 12-hour clock
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {

	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
