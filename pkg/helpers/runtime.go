package helpers

import (
	"fmt"
	"time"
)

func calculateTime(start time.Time) string {
	elapsedTime := float64(time.Since(start).Seconds() * 1000)
	return fmt.Sprintf("%.2f", elapsedTime) + "ms"
}

func QueryTimeTwoOutput[T any](callback func() (T, error)) func() (T, string, error) {
	return func() (T, string, error) {
		start := time.Now()
		result, err := callback()
		return result, calculateTime(start), err
	}
}

func QueryTimeErrorWithParams[T any, P any](callback func(P) error) func(P) (string, error) {
	return func(p P) (string, error) {
		start := time.Now()
		err := callback(p)
		return calculateTime(start), err
	}
}

func QueryTimeOneOutputWithParams[T any, P any](callback func(P) T) func(P) (T, string) {
	return func(p P) (T, string) {
		start := time.Now()
		result := callback(p)
		return result, calculateTime(start)
	}
}

func QueryTimeTwoOutputWithParams[T any, P any](callback func(P) (T, error)) func(P) (T, string, error) {
	return func(p P) (T, string, error) {
		start := time.Now()
		result, err := callback(p)
		return result, calculateTime(start), err
	}
}
