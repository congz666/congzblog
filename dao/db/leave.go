package db

import (
	"fmt"

	"github.com/congz666/congzblog/model"
)

//InsertLeave ...
func InsertLeave(leave *model.Leave) (err error) {
	sqlstr := "insert into `leave`(username,email,content,summary)values(?,?,?,?)"
	_, err = DB.Exec(sqlstr, leave.Username, leave.Email, leave.Content, leave.Summary)
	if err != nil {
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}
	return
}

//GetLeaveList ...
func GetLeaveList() (leaveList []*model.Leave, err error) {
	sqlstr := "select leave_id, username, email, content, summary, create_time from `leave` order by create_time asc"
	err = DB.Select(&leaveList, sqlstr)
	if err != nil {
		fmt.Printf("exec sql:%s failed, err:%v\n", sqlstr, err)
		return
	}
	return
}

//LeaveDelete ...
func LeaveDelete(id int64) {
	_, _ = DB.Exec("delete from `leave` where leave_id = ?", id)
}
