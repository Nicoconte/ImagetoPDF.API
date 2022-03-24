package models

import "time"

type SessionModel struct {
	SessionId string
	CreatedAt time.Time
	UpdatedAt time.Time
}
