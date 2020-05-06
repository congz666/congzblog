package model

//Category ...
//分类
type Category struct {
	CategoryID   int64  `db:"category_id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int    `db:"category_no"` //分类中文章数量

}
