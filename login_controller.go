package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	result := HttpResult{}
	defer func() {
		err := recover()
		if err != nil {
			logger.Println(err)
		}
	}()

	loginReq := &LoginReq{}
	err := c.Bind(loginReq)
	if err != nil {
		result.Fail(-1, "非法请求")
		c.JSON(http.StatusOK, result)
	}

	info := GetSchoolLoginInfo(loginReq.SchoolName)
	loginCookie := ""
	if info["loginType"] == "CLOUD" {
		loginCookie = IAPLogin(loginReq.UserName, loginReq.PassWord, info["loginUrl"], info["loginHost"], info["cookie"])
	} else {
		loginCookie = CASLogin(loginReq.UserName, loginReq.PassWord, info["loginUrl"], info["loginHost"], info["cookie"])
	}
	if loginCookie == "" {
		result.Fail(-2, "登录失败")
		c.JSON(http.StatusOK, result)
	}
	fmt.Println("登录成功:", loginCookie)
	result.Success(loginCookie, info["campusHost"])
	c.JSON(http.StatusOK, result)
}
