package db

import (
	"fmt"

	"github.com/congz666/congzblog/model"
)

//InsertComment ...
func InsertComment(comment *model.Comment) (err error) {
	if comment == nil {
		err = fmt.Errorf("invalid parameter")
		return
	}
	tx, err := DB.Beginx()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()
	sqlstr := `insert 
						into comment(
							content, summary, username, article_id, email					
						)
					values (
							?, ?, ?, ?, ?
					)`
	_, err = tx.Exec(sqlstr, comment.Content, comment.Summary, comment.Username, comment.ArticleID, comment.Email)
	if err != nil {
		return
	}
	sqlstr = `  update 
						article 
					set 
						comment_count = comment_count + 1
					where
						article_id = ?`
	_, err = tx.Exec(sqlstr, comment.ArticleID)
	if err != nil {
		return
	}
	err = tx.Commit()
	return
}

//UpdateViewCount ...
func UpdateViewCount(articleID int64) (err error) {
	sqlstr := ` update 
						article 
					set 
						comment_count = comment_count + 1
					where
						article_id = ?`
	_, err = DB.Exec(sqlstr, articleID)
	if err != nil {
		return
	}
	return
}

//GetCommentList ...
func GetCommentList(articleID int64) (commentList []*model.Comment, err error) {
	sqlstr := `select 
							comment_id, content, username, created_at, summary, email
						from 
							comment 
						where 
							article_id = ?
						order by created_at desc`
	err = DB.Select(&commentList, sqlstr, articleID)
	return
}

//GetAllCommentList ...
func GetAllCommentList() (commentList []*model.Comment, err error) {
	sqlstr := "select comment_id, summary, username,article_id, email from comment  order by article_id asc"
	err = DB.Select(&commentList, sqlstr)
	return
}

//CommentDelete ...
func CommentDelete(id int64) {
	_, _ = DB.Exec("delete from comment where comment_id = ?", id)
}
