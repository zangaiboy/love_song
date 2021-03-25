package main

import (
	"fmt"
	"love_song/middleware"
	"love_song/pkg/setting"
	"love_song/routers"
	"net/http"
)

func main() {
	middleware.InitRedis()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
