package model

import "time"

//Session ...
type Session struct {
	ID        int       `db:"id"`
	UUID      string    `db:"uuid"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
