package model

import (
	"fliteklub/misc"
	"time"
)

type Membership struct {
	misc.GormModel
	User   User      `json:"user"`
	Club   Club      `json:"club"`
	Joined time.Time `json:"joined"`
}
