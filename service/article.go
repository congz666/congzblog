package service

import (
	"github.com/congz666/congzblog/dao/db"
	"github.com/congz666/congzblog/model"
	"github.com/congz666/congzblog/utils/logging"
)

//GetArticleRecordList ...
func GetArticleRecordList() (
	articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleList()
	if err != nil {
		logging.Error(err, "get article list failed")
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	categoryIDs := getCategoryIDs(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIDs)
	if err != nil {
		logging.Error(err, "get category list failed")
		return
	}
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryID := article.CategoryID
		for _, category := range categoryList {
			if categoryID == category.CategoryID {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func getCategoryIDs(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryID := article.CategoryID
		for _, id := range ids {
			if id == categoryID {
				continue LABEL
			}
		}
		ids = append(ids, categoryID)
	}
	return
}

//GetArticleRecordListByID ...
func GetArticleRecordListByID(categoryID int) (
	articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleListByCategoryID(categoryID)
	if err != nil {
		logging.Error(err, "get article list failed")
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	categoryIDs := getCategoryIDs(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIDs)
	if err != nil {
		logging.Error(err, "get category list failed")
		return
	}
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryID := article.CategoryID
		for _, category := range categoryList {
			if categoryID == category.CategoryID {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

//GetArticleDetail ...
func GetArticleDetail(articleID int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(articleID)
	if err != nil {
		logging.Error(err, "get article detail failed")
		return
	}
	category, err := db.GetCategoryByID(articleDetail.ArticleInfo.CategoryID)
	if err != nil {
		logging.Error(err, "get article category failed")
		return
	}
	articleDetail.Category = *category
	return
}

//GetRelativeArticleList ...
func GetRelativeArticleList(articleID int64) (articleList []*model.RelativeArticle, err error) {
	articleList, err = db.GetRelativeArticle(articleID)
	return
}

//GetPrevAndNextArticleInfo ...
func GetPrevAndNextArticleInfo(articleID int64) (prevArticle *model.RelativeArticle,
	nextArticle *model.RelativeArticle, err error) {
	prevArticle, err = db.GetPrevArticleByID(articleID)
	if err != nil {
	}
	nextArticle, err = db.GetNextArticleByID(articleID)
	if err != nil {
	}
	return
}

//InsertArticle ...
func InsertArticle(content, author, title, summary string, categoryID int64) (err error) {
	articleDetail := &model.ArticleDetail{}
	articleDetail.Content = content
	articleDetail.Username = author
	articleDetail.Title = title
	articleDetail.ArticleInfo.CategoryID = categoryID
	articleDetail.Summary = summary
	id, err := db.InsertArticle(articleDetail)
	if err != nil {
		logging.Error(err, "insert article id:%d failed", id)
		return
	}
	return
}

//UpdateArticle ...
func UpdateArticle(content, author, title, summary string, id, categoryID int64) (err error) {
	articleDetail := &model.ArticleDetail{}
	articleDetail.ID = id
	articleDetail.Content = content
	articleDetail.Username = author
	articleDetail.Title = title
	articleDetail.ArticleInfo.CategoryID = categoryID
	articleDetail.Summary = summary
	err = db.UpdateArticle(articleDetail)
	if err != nil {
		logging.Error(err, "update article failed")
		return
	}
	return
}

//DeliverArticleID ...
func DeliverArticleID(id int64) {
	err := db.ArticleDelete(id)
	if err != nil {
		logging.Error(err)
	}
}

//ArticleCount ...
func ArticleCount() (count int) {
	count, err := db.ArticleCount()
	if err != nil {
		logging.Error(err)
	}
	return
}
