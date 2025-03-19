package models

type Session struct {
	ID     string `json:"id" db:"id"`
	Active bool   `json:"active" db:"active"`
}
