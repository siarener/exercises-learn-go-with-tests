// Package clockface provides functions that calculate the positions of the hands
// of an analogue clock.
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

// A Point is a Cartesian coordinate.
// They are used in the package to represent the unit vecotr from the origin of a
// clock hand.
type Point struct {
	X float64
	Y float64
}

// SecondsInRadians returns the angle of the second hand from 12 o'clock in radians.
func SecondsInRadians(t time.Time) float64 {
	/* 1 second is 2pi/60 radians
	-> float64(t.Second()) * (math.Pi / 30)
	-> re-arrange to avoid first dividing down and then multiplying up,
		as this would results in a very inaccurate result because of
		inaccurate floating point arithmetic
	*/
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

// SecondHandPoint is the unit vector of the second hand at time `t`,
// represented as a Point
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

// MinutesInRadians returns the angle of the minute hand from 12 o'clock in radians.
func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / (float64(t.Minute()))))
}

// MinuteHandPoint is the unit vector of the minute hand at time `t`,
// represented as a Point.
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

// HoursInRadians returns the angle of the hour hand from 12 o'clock in radians.
func HoursInRadians(t time.Time) float64 {
	// get the remainder of the hour divided by 12, as this is a 12-hour clock
	return (MinutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

// HourHandPoint is the unit vecotr of the hour hand ad time `t`,
// represented as a Point.
func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {

	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
