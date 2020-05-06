package db

import (
	"database/sql"
	"time"

	"github.com/congz666/congzblog/model"
	"github.com/congz666/congzblog/utils/logging"
)

//CreateSession ...
func CreateSession(name string) (session model.Session) {
	sqlstr := "insert into session (uuid,name,created_at) values(?,?,?)"
	_, err := DB.Exec(sqlstr, createUUID(), name, time.Now())
	if err != nil {
		logging.Error(err, "create session failed")
		return
	}
	sqlstr = "select id,uuid,name,created_at from session"
	err = DB.QueryRow(sqlstr).Scan(&session.ID, &session.UUID, &session.Name, &session.CreatedAt)
	if err != nil {
		logging.Error(err, "select session failed")
		return
	}

	return
}

//UpdateSession ...
func UpdateSession(name string) (session model.Session) {
	sqlstr := "update session set uuid = ?, created_at = ? where name = ?"
	_, err := DB.Exec(sqlstr, createUUID(), time.Now(), name)
	if err != nil {
		logging.Error(err, "update session failed")
		return
	}
	sqlstr = "select id,uuid,name,created_at from session"
	err = DB.QueryRow(sqlstr).Scan(&session.ID, &session.UUID, &session.Name, &session.CreatedAt)
	if err != nil {
		logging.Error(err, "select session failed")
		return
	}
	return
}

//Check ...
func Check(UUID string) (valid bool, err error) {
	var sess model.Session
	err = DB.QueryRow("select id, uuid, name, created_at from session where uuid = ?", UUID).Scan(&sess.ID, &sess.UUID, &sess.Name, &sess.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return
}

//DeleteAllSession ...
func DeleteAllSession() (err error) {
	sqlstr := "delete  from session"
	_, err = DB.Exec(sqlstr)
	if err != nil {
		logging.Error("delete session failed")
	}
	return
}

//IsSessionExist ...
func IsSessionExist(name string) (exists bool) {
	var sessname string
	err := DB.QueryRow("select name from session where name = ?", name).Scan(&sessname)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		logging.Error(err)
		return
	}
	exists = true
	return
}
