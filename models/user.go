package models

import (
	"fmt"
)

type User struct {
	ID int8 `json:"id"`
	//Nickname   string    `json:"nickname"`
	//Username   string    `json:"username"`
	//Password   string    `json:"password"`
	//HeadUrl    string    `json:"head_url"`
	//State      int       `json:"state"`
	//CreateTime time.Time `json:"create_time"`
	//UpdateTime time.Time `json:"update_time"`
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (Users []User) {
	rows, err := db.Query("SELECT * FROM users LIMIT 10 OFFSET 0")
	checkErr(err)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		checkErr(err)
		fmt.Println(id)
	}
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
