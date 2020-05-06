package model

import (
	"time"
)

//Comment ...
//评论
type Comment struct {
	ID         int64     `db:"comment_id"`
	Content    string    `db:"content"`
	Summary    string    `db:"summary"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Status     int       `db:"status"`
	ArticleID  int64     `db:"article_id"`
}
