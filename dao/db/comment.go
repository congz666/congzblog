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
							content, summary, username, article_id					
						)
					values (
							?, ?, ?,?
					)`
	_, err = tx.Exec(sqlstr, comment.Content, comment.Summary, comment.Username, comment.ArticleID)
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
func GetCommentList(articleID int64, pageNum, pageSize int) (commentList []*model.Comment, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
							comment_id, content, username, create_time
						from 
							comment 
						where 
							article_id = ? and 
							status = 1
						order by create_time desc
						limit ?, ?`
	err = DB.Select(&commentList, sqlstr, articleID, pageNum, pageSize)
	return
}

//GetAllCommentList ...
func GetAllCommentList() (commentList []*model.Comment, err error) {
	sqlstr := "select comment_id, summary, username,article_id from comment order by comment_id asc"
	err = DB.Select(&commentList, sqlstr)
	return
}

//CommentDelete ...
func CommentDelete(id int64) {
	_, _ = DB.Exec("delete from comment where comment_id = ?", id)
}
