package internal

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
