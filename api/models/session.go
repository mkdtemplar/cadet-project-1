package models

import (
	"time"
)

var Sessions = map[string]Session{}

type Session struct {
	Expiry time.Time
}

func (s *Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
