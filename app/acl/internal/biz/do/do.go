package do

import (
	"time"
)

type Role struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	IsSystem    bool
}
