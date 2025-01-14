package helpers

import (
	"fmt"
	"time"
)

func calculateTime(start time.Time) string {
	end := time.Now()
	elapsedTime := float64(end.Sub(start).Seconds() * 1000)
	return fmt.Sprintf("%.2f", elapsedTime) + "ms"
}

func QueryTimeTwoOutput[T any](callback func() (T, error)) func() (T, string, error) {
	return func() (T, string, error) {
		start := time.Now()
		result, err := callback()
		return result, calculateTime(start), err
	}
}

func QueryTimeErrorWithOneParam[T any, P any](callback func(P) error) func(P) (string, error) {
	return func(p P) (string, error) {
		start := time.Now()
		err := callback(p)
		return calculateTime(start), err
	}
}

func QueryTimeErrorWithTwoParams[T any, P1, P2 any](callback func(P1, P2) error) func(P1, P2) (string, error) {
	return func(p1 P1, p2 P2) (string, error) {
		start := time.Now()
		err := callback(p1, p2)
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
