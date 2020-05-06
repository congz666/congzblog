package db

import (
	"database/sql"
	"fmt"

	"github.com/congz666/congzblog/model"
)

//InsertArticle ...
func InsertArticle(article *model.ArticleDetail) (articleID int64, err error) {
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlstr := `insert into 
					article(content, summary, title, username, 
						category_id, view_count, comment_count)
				values(?, ?, ?, ?, ?, ?, ?)`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryID,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articleID, err = result.LastInsertId()
	_, err = DB.Exec("update category set category_no=category_no+1 where category_id=?", article.ArticleInfo.CategoryID)
	if err != nil {
		fmt.Println("update failed")
		return
	}
	return
}

//GetArticleList ...
//status=1
func GetArticleList() (articleList []*model.ArticleInfo, err error) {
	sqlstr := `select 
						article_id, summary, title, view_count,
						 created_at, updated_at, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
					order by created_at desc`
	err = DB.Select(&articleList, sqlstr)
	return
}

//GetRecycleList ...
//status=0
func GetRecycleList() (articleList []*model.ArticleInfo, err error) {
	sqlstr := `select 
						article_id, summary, title, view_count,
						 created_at, updated_at, comment_count, username, category_id
					from 
						article 
					where 
						status = 0
					order by created_at desc
					`
	err = DB.Select(&articleList, sqlstr)
	return
}

//GetArticleDetail ...
func GetArticleDetail(articleID int64) (articleDetail *model.ArticleDetail, err error) {
	if articleID < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", articleID)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select 
							article_id, summary, title, content,
							 created_at, updated_at, username, category_id, 
							 view_count , comment_count
						from 
							article 
						where 
							article_id = ?
						`
	err = DB.Get(articleDetail, sqlstr, articleID)
	articleDetail.ArticleInfo.CategoryID = articleDetail.Category.CategoryID
	return
}

//GetArticleListByCategoryID ...
func GetArticleListByCategoryID(categoryID int) (articleList []*model.ArticleInfo, err error) {
	sqlstr := `select 
						article_id, summary, title, view_count,
						 created_at, updated_at, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
						and
						category_id = ?
					order by created_at desc`
	err = DB.Select(&articleList, sqlstr, categoryID)
	return
}

//GetRelativeArticle ...
func GetRelativeArticle(articleID int64) (articleList []*model.RelativeArticle, err error) {
	var categoryID int64
	sqlstr := "select category_id from article where article_id=?"
	err = DB.Get(&categoryID, sqlstr, articleID)
	if err != nil {
		return
	}
	sqlstr = "select article_id, title from article where category_id=? and article_id !=?  limit 2"
	err = DB.Select(&articleList, sqlstr, categoryID, articleID)
	return
}

//GetPrevArticleByID ...
func GetPrevArticleByID(articleID int64) (info *model.RelativeArticle, err error) {

	info = &model.RelativeArticle{
		ArticleID: -1,
	}
	sqlstr := "select article_id, title from article where article_id < ? order by article_id desc limit 1"
	err = DB.Get(info, sqlstr, articleID)
	if err != nil {
		return
	}
	return
}

//GetNextArticleByID ...
func GetNextArticleByID(articleID int64) (info *model.RelativeArticle, err error) {
	info = &model.RelativeArticle{
		ArticleID: -1,
	}
	sqlstr := "select article_id, title from article where article_id > ? order by article_id asc limit 1"
	err = DB.Get(info, sqlstr, articleID)
	if err != nil {
		return
	}
	return
}

//IsArticleExist ...
func IsArticleExist(articleID int64) (exists bool, err error) {
	var id int64
	sqlstr := "select article_id from article where article_id=?"
	err = DB.Get(&id, sqlstr, articleID)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return
}

//ArticleDelete ...
func ArticleDelete(id int64) (err error) {
	sqlstr := "update article set status=0 where article_id=?"
	_, err = DB.Exec(sqlstr, id)
	if err != nil {
		return
	}

	//分类文章减一
	_, err = DB.Exec("update category set category_no=category_no-1 where category_id=(select category_id from article where article_id=?) ", id)
	if err != nil {
		return
	}

	return
}

//UpdateArticle ...
func UpdateArticle(article *model.ArticleDetail) (err error) {
	_, err = DB.Exec("update category set category_no=category_no-1 where category_id=(select category_id from article where article_id=?)", article.ID)
	if err != nil {
		return
	}
	sqlstr := "update article set title=?, username=?, category_id=?, summary=?,content=? where article_id=? "
	_, err = DB.Exec(sqlstr, article.Title, article.Username, article.ArticleInfo.CategoryID, article.Summary, article.Content, article.ID)
	if err != nil {
		return
	}
	_, err = DB.Exec("update category set category_no=category_no+1 where category_id=?", article.ArticleInfo.CategoryID)
	if err != nil {
		return
	}
	return
}

//ArticleCount ...
func ArticleCount() (count int, err error) {
	sqlstr := "select count(*) from article"
	err = DB.Get(&count, sqlstr)
	if err != nil {
		return
	}
	return
}

//RecycleResume ...
func RecycleResume(id int64) (err error) {
	sqlstr := "update article set status=1 where article_id=?"
	_, err = DB.Exec(sqlstr, id)
	if err != nil {
		return
	}
	//分类文章加一
	_, err = DB.Exec("update category set category_no=category_no+1 where category_id=(select category_id from article where article_id=?) ", id)
	if err != nil {
		return
	}
	return
}

//RecycleDelete ...
func RecycleDelete(id int64) {
	_, _ = DB.Exec("delete from article where article_id = ?", id)
}
