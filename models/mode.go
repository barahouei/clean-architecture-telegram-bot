// Package models contain entities that are essential for the app.
package models

// Mode shows whether the application is in development mode or production mode.
type Mode uint

const (
	_ Mode = iota
	Development
	Production
)
