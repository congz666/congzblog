package model

import (
	"time"
)

//Leave ...
//留言
type Leave struct {
	ID         int64     `db:"leave_id"`
	Content    string    `db:"content"`
	Summary    string    `db:"summary"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Email      string    `db:"email"`
}
