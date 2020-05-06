package service

import (
	"github.com/congz666/congzblog/dao/db"
	"github.com/congz666/congzblog/model"
	"github.com/congz666/congzblog/utils/logging"
)

//GetRecycleRecordList ...
func GetRecycleRecordList() (articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetRecycleList()
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

//DeliverRecycleID ...
func DeliverRecycleID(id int64) {
	db.RecycleDelete(id)
}

//RecycleResume ...
func RecycleResume(id int64) {
	err := db.RecycleResume(id)
	if err != nil {
		logging.Error(err)
	}
}
