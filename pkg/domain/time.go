package domain

import "time"

// CreatedAt returns the current date and time
func CreatedAt() time.Time {
	utc := time.Now()
	return utc
}
