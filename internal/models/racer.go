package models

import "time"

type Racer struct {
	ID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string `sql:"type:varchar(100)"`
	Image     string `sql:"type:varchar(500)"`
	Rating    int64  `sql:"type:bigint"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
