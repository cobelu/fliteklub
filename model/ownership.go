package model

import (
	"fliteklub/misc"
	"time"
)

type Ownership struct {
	misc.GormModel
	Club     Club      `json:"club"`
	Aircraft Aircraft  `json:"aircraft"`
	Added    time.Time `json:"added"`
}
