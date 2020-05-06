package db

import (
	"github.com/congz666/congzblog/model"

	"github.com/jmoiron/sqlx"
)

//InsertCategory ...
func InsertCategory(category *model.Category) (categoryID int64, err error) {
	sqlstr := "insert into category(category_name, category_no)values(?,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryID, err = result.LastInsertId()
	return
}

//GetCategoryByID ...
func GetCategoryByID(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlstr := "select category_id, category_name, category_no from category where category_id=?"
	err = DB.Get(category, sqlstr, id)
	return
}

//GetCategoryList ...
func GetCategoryList(categoryIDs []int64) (categoryList []*model.Category, err error) {
	sqlstr, args, err := sqlx.In("select category_id, category_name, "+
		"category_no from category where category_id in(?)", categoryIDs)
	if err != nil {
		return
	}
	err = DB.Select(&categoryList, sqlstr, args...)
	return
}

//GetAllCategoryList ...
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlstr := "select category_id, category_name, category_no from category order by category_no desc"
	err = DB.Select(&categoryList, sqlstr)
	return
}

//CategoryDelete ...
func CategoryDelete(id int64) {
	_, _ = DB.Exec("delete from category where category_id = ?", id)
}

//UpdateCategory ...
func UpdateCategory(category *model.Category) (err error) {
	sqlstr := "update category set category_name=?,category_no=? where category_id=?"
	_, err = DB.Exec(sqlstr, category.CategoryName, category.CategoryNo, category.CategoryID)
	if err != nil {
		return
	}
	return
}

//CategoryCount ...
func CategoryCount() (count int, err error) {
	sqlstr := "select count(*) from category"
	err = DB.Get(&count, sqlstr)
	if err != nil {
		return
	}
	return
}
