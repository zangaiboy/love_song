package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"love_song/models"
	"love_song/pkg/e"
	"love_song/pkg/setting"
	"love_song/pkg/util"
	"net/http"
)

//获取用户列表
func GetUsers(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS

	data["lists"] = models.GetUsers(util.GetPage(c), setting.PageSize, maps)
	//data["total"] = models.GetUserTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增用户
//func AddUser(c *gin.Context) {
//	nickName := c.Query("nickname")
//	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
//
//	valid := validation.Validation{}
//	valid.Required(nickName, "nickname").Message("名称不能为空")
//	valid.MaxSize(nickName, 100, "nickname").Message("名称最长为100字符")
//	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
//
//	code := e.INVALID_PARAMS
//	if !valid.HasErrors() {
//		if !models.ExistUserByName(nickName) {
//			code = e.SUCCESS
//			models.AddUser(nickName, state)
//		} else {
//			code = e.ERROR_EXIST_TAG
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
//
////修改用户
//func EditUser(c *gin.Context) {
//
//	id := com.StrTo(c.Param("id")).MustInt()
//	nickName := c.Query("nickname")
//
//	valid := validation.Validation{}
//
//	var state int = -1
//	if arg := c.Query("state"); arg != "" {
//		state = com.StrTo(arg).MustInt()
//		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
//	}
//	valid.Required(id, "id").Message("ID不能为空")
//	valid.MaxSize(nickName, 100, "nickname").Message("名称最长为100字符")
//
//	code := e.INVALID_PARAMS
//	if !valid.HasErrors() {
//		code = e.SUCCESS
//		if models.ExistUserByID(id) {
//			data := make(map[string]interface{})
//			if nickName != "" {
//				data["nickname"] = nickName
//			}
//			if state != -1 {
//				data["state"] = state
//			}
//
//			models.EditUser(id, data)
//		} else {
//			code = e.ERROR_NOT_EXIST_TAG
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
//
////删除用户
//func DeleteUser(c *gin.Context) {
//	id := com.StrTo(c.Param("id")).MustInt()
//
//	valid := validation.Validation{}
//	valid.Min(id, 1, "id").Message("ID必须大于0")
//
//	code := e.INVALID_PARAMS
//	if !valid.HasErrors() {
//		code = e.SUCCESS
//		if models.ExistUserByID(id) {
//			models.DeleteUser(id)
//		} else {
//			code = e.ERROR_NOT_EXIST_TAG
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
