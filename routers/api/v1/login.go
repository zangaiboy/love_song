package v1

import (
	"github.com/gin-gonic/gin"
)

//登陆
func Login(c *gin.Context) {
	//loginDto := dto.LoginDto{}
	//c.BindJSON(&loginDto)
	//fmt.Println("---body/--- \r\n " + loginDto.Username)
	//user := models.GetUserByUsername(loginDto.Username)
	//if loginDto.Password == user.Password {
	//	code := e.SUCCESS
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": code,
	//		"msg":  e.GetMsg(code),
	//	})
	//	return
	//}
	//code := e.ERROR
	//c.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//})
}
