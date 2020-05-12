package service

import (
	"time"

	"github.com/congz666/congzblog/dao/db"
	"github.com/congz666/congzblog/model"
	"github.com/congz666/congzblog/utils/logging"
)

//InsertComment ...
func InsertComment(comment, author, email, summary string, articleID int64) (err error) {
	exist, err := db.IsArticleExist(articleID)
	if err != nil {
		logging.Error(err, "query database failed")
		return
	}
	if exist == false {
		logging.Error(err, "article id:%d not found", articleID)
		return
	}
	var c model.Comment
	c.ArticleID = articleID
	c.Content = comment
	c.Summary = summary
	c.Username = author
	c.CreatedAt = time.Now()
	c.Email = email
	err = db.InsertComment(&c)
	return
}

//GetCommentList ...
func GetCommentList(articleID int64) (commentList []*model.Comment, err error) {
	exist, err := db.IsArticleExist(articleID)
	if err != nil {
		logging.Error(err, "query database failed")
		return
	}
	if exist == false {
		logging.Error(err, "article id:%d not found", articleID)
		return
	}
	commentList, err = db.GetCommentList(articleID)
	return
}

//GetAllCommentList ...
func GetAllCommentList() (commentList []*model.Comment, err error) {
	commentList, err = db.GetAllCommentList()
	if err != nil {
		logging.Error(err, "get comment list failed")
		return
	}
	return
}

//DeliverCommentID ...
func DeliverCommentID(id int64) {
	db.CommentDelete(id)
}
