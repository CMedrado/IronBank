package domain

import "time"

// CreatedAt returns the current date and time
func CreatedAt() string {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	utc := time.Now().In(loc).Format("02/01/2006 15:04:05")
	return utc
}
