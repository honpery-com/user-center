package model

import "time"

// UserStatus is int value
type UserStatus = int

const (
	// UserStatusActive = 0
	UserStatusActive UserStatus = iota

	// UserStatusDisable = 1
	UserStatusDisable
)

// User table struct.
type User struct {
	ID       string     `json:"_id"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	CreateAt time.Timer `json:"create_at"`
	UpdateAt time.Timer `json:"update_at"`
}
