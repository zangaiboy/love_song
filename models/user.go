package models

import (
	"time"
)

type User struct {
	ID         int8      `json:"id"`
	Nickname   string    `json:"nickname"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	HeadUrl    string    `json:"head_url"`
	State      int       `json:"state"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (Users []User) {
	rows, err := db.Query("SELECT * FROM users LIMIT 10 OFFSET 0")
	checkErr(err)
	for rows.Next() {
		var id int8
		var nickname, username, password, headUrl string
		var state int
		var createTime, updateTime time.Time
		err = rows.Scan(&id, &nickname, &username, &password, &headUrl, &state, &createTime, &updateTime)
		checkErr(err)
		Users = append(Users, User{id, nickname, username, password, headUrl, state, createTime, updateTime})
	}
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
