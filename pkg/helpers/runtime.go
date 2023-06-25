package helpers

import (
	"time"
	"fmt"
)

func calculateTime(start time.Time) {
	elapsedTime := float64(time.Since(start).Seconds() * 1000)
	OutputTime(fmt.Sprintf("%.2f", elapsedTime) + "ms")
}

func QueryTimeTwoOutput[T any](callback func()(T, error)) func() (T, error) {
	start := time.Now()

	return func() (T, error) {
		defer calculateTime(start)
		return callback()
	}
}

func QueryTimeTwoOutputWithParams[T any, P any](callback func(P)(T, error)) func(P) (T, error) {
	start := time.Now()

	return func(p P) (T, error) {
		defer calculateTime(start)
		return callback(p)
	}
}

func QueryTimeOneOutputWithParams[T any, P any](callback func(P)(T)) func(P) (T) {
	start := time.Now()

	return func(p P) (T) {
		defer calculateTime(start)
		return callback(p)
	}
}
