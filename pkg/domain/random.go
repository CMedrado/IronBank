package domain

import (
	"github.com/google/uuid"
)

// Random returns a random number
func Random() (uuid.UUID, error) {
	return uuid.NewRandom()
}
