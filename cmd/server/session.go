package main

import (
	"time"
)

type Session struct {
	ID        	int
	SessionId 	string
	UserId		int
	CreatedAt 	time.Time
	ModifiedAt	time.Time
}

