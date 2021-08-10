package domain

import "time"

// CreatedAt returns the current date and time
func CreatedAt() time.Time {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	utc := time.Now().In(loc)
	return utc
}
