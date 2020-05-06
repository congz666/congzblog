package model

import "time"

//ArticleInfo ...
type ArticleInfo struct {
	ID           int64     `db:"article_id"`
	CategoryID   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

//ArticleDetail ...
type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

//ArticleRecord ...
type ArticleRecord struct {
	ArticleInfo
	Category
}
