package models

import (
	"net/http"
	"sync"
	"time"
)

var mu = sync.Mutex{}
var sessions = map[string]Session{}

func AddSession(key string, session Session) {
	mu.Lock()
	sessions[key] = session
	mu.Unlock()
}

func GetSession() map[string]Session {
	return sessions
}

type Session struct {
	Expiry time.Time
}

var Cookie http.Cookie

func (s *Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
