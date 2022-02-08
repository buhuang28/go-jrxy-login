package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			logger.Println(err)
		}
	}()
	result := HttpResult{}

	loginReq := &LoginReq{}
	err := c.Bind(loginReq)
	if err != nil {
		result.Fail(-1, "非法请求")
		c.JSON(http.StatusOK, result)
	}

	info := GetSchoolLoginInfo(loginReq.SchoolName)
	loginCookie := CASLogin(loginReq.UserName, loginReq.PassWord, info["loginUrl"], "", info["cookie"])
	if loginCookie == "" {
		result.Fail(-2, "登录失败")
		c.JSON(http.StatusOK, result)
	}
	fmt.Println("登录成功:", loginCookie)
	result.Success(loginCookie, info["campusHost"])
	c.JSON(http.StatusOK, result)
}
