// Package models contain entities that are essential for the app.
package models

// Language shows in which language a user sees menus and messages.
type Language uint

const (
	_ Language = iota
	En
	Fa
)
