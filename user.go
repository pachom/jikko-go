// user.go
package main

import "time"

type User struct {
	ID        int       `json:"id,omitempty"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
