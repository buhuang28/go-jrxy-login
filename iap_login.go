package main

import (
	"github.com/valyala/fastjson"
)

func IAPLogin(userName, passWord, loginUrl, loginHost, cookie string) string {
	defer func() {
		e := recover()
		if e != nil {
			logger.Println(e)
		}
	}()

	ck := CookieStr2Map(cookie)
	success, resp := PostJson(loginHost+"iap/security/lt", &ck, nil, nil)
	if !success {
		logger.Println("IAP登录获取lt失败")
		return ""
	}

	bytes, err := fastjson.ParseBytes(resp.Data)
	if err != nil {
		logger.Println("IAP登录解析lt json失败:", string(resp.Data))
		return ""
	}

	lt := string(bytes.Get("result").GetStringBytes("_lt"))
	params := make(map[string]string)
	params["lt"] = lt
	params["rememberMe"] = "false"
	params["dllt"] = ""
	params["mobile"] = ""
	params["username"] = userName
	params["password"] = passWord

	loginSuccess, httpResp := PostForm(loginHost+"iap/doLogin", &ck, nil, params)
	if !loginSuccess {
		logger.Println("iap登录失败:", string(httpResp.Data))
		return ""
	}
	ts := 0
	for httpResp.Localtion != "" && ts < 10 {
		ts++
		loginSuccess, httpResp = PostForm(httpResp.Localtion, &ck, nil, nil)
	}
	if !loginSuccess {
		logger.Println("iap登录失败2:", string(httpResp.Data))
		return ""
	}
	return CookieMap2Str(*httpResp.Cookie)
}
