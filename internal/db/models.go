// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Todo struct {
	ID        int64
	Task      string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}