package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type LoginAttempt struct {
	bun.BaseModel `bun:"public.login_attempts"`
	Id            int64     `json:"id" bun:"id,pk,autoincrement"`
	Email         string    `json:"email" bun:"email,notnull"`
	AttemptTime   time.Time `json:"attempt_time" bun:"attempt_time,notnull"`
	IPAddress     string    `json:"ip_address" bun:"ip_address,notnull"`
	UserAgent     string    `json:"user_agent" bun:"user_agent"`
	Device        string    `json:"device" bun:"device,nullzero"`
	Os            string    `json:"os" bun:"os,nullzero"`
	OsVersion     string    `json:"os_version" bun:"os_version,nullzero"`
	Successful    bool      `bun:"successful,notnull"`
	Reason        string    `bun:"reason,nullzero"`
}
