// Package models contain entities that are essential for the app.
package models

import "time"

// User represents user table structure in a database.
type User struct {
	ID         int
	TelegramID int64
	Username   string
	FirstName  string
	LastName   string
	JoinedAt   time.Time
	Language   Language
}
