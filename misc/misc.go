package misc

import "gorm.io/gorm"
import "time"

type GormModel gorm.Model

func now() time.Time {
	utc := time.Now().UTC()
	return utc
}
