package middleware

import (
	"fmt"
	"gopkg.in/redis.v5"
	"love_song/pkg/setting"
)

func InitRedis() {
	sec, err := setting.Cfg.GetSection("redis")
	host := sec.Key("HOST").String()
	password := sec.Key("PASSWORD").String()
	db, _ := sec.Key("DB").Int()

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println("redis connection ...", pong, err)
}
