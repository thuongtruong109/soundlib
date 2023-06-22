package helpers

import (
	"time"
	"math/rand"
)

func GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000)
}