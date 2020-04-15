package model

import (
	"time"
)

type Phone struct {
	PhoneNumber      string
	ActivationDate   time.Time
	DeactivationDate time.Time
}
