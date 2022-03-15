package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func CASLogin(userName, passWord, loginUrl, loginHost, cookie string) string {
	ck := CookieStr2Map(cookie)
	success, resp := GetReqeust(loginUrl, &ck, nil, nil)
	if !success {
		fmt.Println("登录失败，访问", loginUrl, "失败")
		return ""
	}
	pwSalt := ""
	loginType := 0
	params := make(map[string]string)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Data))
	if err != nil {
		fmt.Println("登录失败，解析HTML失败")
		return ""
	}
	form := doc.Find("#casLoginForm")
	if form != nil && len(form.Nodes) != 0 {
		loginType = 0
	} else if form = doc.Find("#loginFromId"); form != nil && len(form.Nodes) != 0 {
		loginType = 1
	} else if form = doc.Find("#fm1"); form != nil && len(form.Nodes) != 0 {
		loginType = 2
	}
	if form == nil {
		fmt.Println("登录失败，找不到登录表单")
		return ""
	}

	form.Each(func(i int, s *goquery.Selection) {
		if (loginType != 0 && i == 1) || (loginType == 0 && i == 0) {
			s.Find("input").Each(func(j int, s2 *goquery.Selection) {
				for _, v := range s2.Nodes {
					key := ""
					value := ""
					for _, v2 := range v.Attr {
						if v2.Val == "rememberMe" {
							continue
						}
						if v2.Key == "name" {
							key = v2.Val
						}
						if v2.Key == "value" {
							value = v2.Val
						}
					}
					if key != "" {
						params[key] = value
					} else {
						if value != "" {
							pwSalt = value
						}
					}
				}
			})
		}
	})
	params["username"] = userName

	if pwSalt == "" {
		pattern := regexp.MustCompile(`\"(\w{16})\"`)
		all := pattern.FindAll(resp.Data, -1)
		if len(all) == 1 {
			pwSalt = string(all[0])
		}
	}
	pwSalt = strings.ReplaceAll(pwSalt, `"`, "")

	switch loginType {
	case 0, 1:
		//AES
		if pwSalt != "" {
			params["password"], _ = AESEncrypt([]byte(RandStr(64)+passWord), []byte(pwSalt))
		} else {
			params["password"] = passWord
		}
		success, httpResp := PostForm(loginUrl, &ck, nil, params)
		if !success {
			fmt.Println("登录失败，可能是密码错误1")
			return ""
		}
		if httpResp.Localtion == "" {
			fmt.Println("登录失败，可能是密码错误2")
			return ""
		}
		success, httpResp = PostForm(httpResp.Localtion, &ck, nil, nil)
		if !success {
			fmt.Println("登录失败，可能是密码错误3")
			return ""
		} else {
			return CookieMap2Str(*httpResp.Cookie)
		}
	case 2:
		//RSA,懒得做...告辞
	}
	fmt.Println("我也母鸡啥问题啊...")
	return ""
}
