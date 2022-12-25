package models

import (
	"net/http"
	"time"
)

var Sessions = map[string]Session{}

type Session struct {
	Expiry time.Time
}

var Cookie *http.Cookie

func (s *Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
