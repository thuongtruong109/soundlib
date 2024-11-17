package helpers

import (
	"fmt"
	"time"
)

func calculateTime(start time.Time) {
	elapsedTime := float64(time.Since(start).Seconds() * 1000)
	OutputTime(fmt.Sprintf("%.2f", elapsedTime) + "ms")
}

func calculateTime2(start time.Time) string {
	elapsedTime := float64(time.Since(start).Seconds() * 1000)
	return fmt.Sprintf("%.2f", elapsedTime) + "ms"
}

func QueryTimeTwoOutput[T any](callback func() (T, error)) func() (T, error) {
	start := time.Now()

	return func() (T, error) {
		defer calculateTime(start)
		return callback()
	}
}

func QueryTimeTwoOutput2[T any](callback func() (T, error)) func() (T, string, error) {
	start := time.Now()

	return func() (T, string, error) {
		result, err := callback()
		return result, calculateTime2(start), err
	}
}

func QueryTimeTwoOutputWithParams[T any, P any](callback func(P) (T, error)) func(P) (T, error) {
	start := time.Now()

	return func(p P) (T, error) {
		defer calculateTime(start)
		return callback(p)
	}
}

func QueryTimeOneOutputWithParams[T any, P any](callback func(P) T) func(P) T {
	start := time.Now()

	return func(p P) T {
		defer calculateTime(start)
		return callback(p)
	}
}
