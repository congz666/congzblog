package model

//RelativeArticle ...
//关系文章
type RelativeArticle struct {
	ArticleID int64  `db:"article_id"`
	Title     string `db:"title"`
}
