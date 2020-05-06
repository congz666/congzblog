package service

import (
	"github.com/congz666/congzblog/dao/db"
	"github.com/congz666/congzblog/model"
	"github.com/congz666/congzblog/utils/logging"
)

//InsertLeave ...
func InsertLeave(username, email, content, summary string) (err error) {
	var leave model.Leave
	leave.Content = content
	leave.Summary = summary
	leave.Email = email
	leave.Username = username
	err = db.InsertLeave(&leave)
	if err != nil {
		logging.Error(err, "insert leave failed")
		return
	}
	return
}

//GetLeaveList ...
func GetLeaveList() (leaveList []*model.Leave, err error) {
	leaveList, err = db.GetLeaveList()
	if err != nil {
		logging.Error(err, "get leave list failed")
		return
	}
	return
}

//DeliverLeaveID ...
func DeliverLeaveID(id int64) {
	db.LeaveDelete(id)
}
