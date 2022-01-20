package user

import (
	"blackboard/model"

	"github.com/gin-gonic/gin"
)

//@Summary 登录
//@Tags user
//@Description 一站式登录
//@Accep applica/json
//@Produce applic/json
//@Param object body model.User true "登录用户信息"
//Success 200 {object} Token "将student_id作为token保留"
// @Success 200 {object} handler.Response "{"msg":"将student_id作为token保留"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Password or account wrong."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var user model.User
	//
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if user.StudentID == "" {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	pwd := user.PassWord
	//首次登录，验证一站式

}
