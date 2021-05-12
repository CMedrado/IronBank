package domain

import (
	"math/rand"
	"time"
)

func CreatedAt() string {
	return time.Now().Format("02/01/2006 03:03:05")
}

func Random() int {
	return rand.Intn(100000000)
}
