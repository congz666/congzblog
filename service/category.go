package service

import (
	"github.com/congz666/congzblog/dao/db"
	"github.com/congz666/congzblog/model"
	"github.com/congz666/congzblog/utils/logging"
)

//GetAllCategoryList ...
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		logging.Error(err, "get category list failed")
		return
	}
	return
}

//GetCategory ...
func GetCategory(id int64) (category *model.Category, err error) {
	category, err = db.GetCategoryByID(id)
	if err != nil {
		logging.Error(err, "get category by id failed")
		return
	}
	return
}

//InsertCategory ...
func InsertCategory(categoryname string, categoryno int) (err error) {
	CategoryDetail := &model.Category{}
	CategoryDetail.CategoryName = categoryname
	CategoryDetail.CategoryNo = categoryno
	id, err := db.InsertCategory(CategoryDetail)
	if err != nil {
		logging.Error(err, "insert category id:%d failed", id)
		return
	}
	return
}

//DeliverCategoryID ...
func DeliverCategoryID(id int64) {
	db.CategoryDelete(id)
}

//UpdateCategory ...
func UpdateCategory(categoryname string, categoryno int, categoryid int64) (err error) {
	CategoryDetail := &model.Category{}
	CategoryDetail.CategoryName = categoryname
	CategoryDetail.CategoryNo = categoryno
	CategoryDetail.CategoryID = categoryid
	err = db.UpdateCategory(CategoryDetail)
	if err != nil {
		logging.Error(err, "update category failed")
		return
	}
	return
}

//CategoryCount ...
func CategoryCount() (count int) {
	count, err := db.CategoryCount()
	if err != nil {
		logging.Error(err)
	}
	return
}
