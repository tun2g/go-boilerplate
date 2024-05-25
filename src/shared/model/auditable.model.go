package model

import (
	"time"
)

type Auditable struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt time.Time
}