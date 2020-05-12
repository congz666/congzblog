package model

import (
	"time"
)

//Comment ...
//评论
type Comment struct {
	ID        int64     `db:"comment_id"`
	Content   string    `db:"content"`
	Summary   string    `db:"summary"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	Email     string    `db:"email"`
	ArticleID int64     `db:"article_id"`
}
