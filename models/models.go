package models

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
	"love_song/pkg/setting"
)

var db *sql.DB

func init() {
	var (
		err                                        error
		dbType, dbName, user, password, host, port string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DBNAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()
	db, err = sql.Open(dbType, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName))
	if err != nil {
		log.Println("db err", err)
	}
}
