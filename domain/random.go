package domain

import "math/rand"

// Random returns a random number
func Random() int {
	return rand.Intn(100000000)
}
