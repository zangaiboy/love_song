package routers

import (
	"github.com/gin-gonic/gin"

	"love_song/pkg/setting"
	"love_song/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取用户列表
		apiv1.GET("/user", v1.GetUsers)
		//新增用户
		apiv1.POST("/user", v1.AddUser)
		//修改用户
		apiv1.PUT("/user/:id", v1.EditUser)
		//删除用户
		apiv1.DELETE("/user/:id", v1.DeleteUser)
	}

	return r
}
