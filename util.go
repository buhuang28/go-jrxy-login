package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	HostReg = regexp.MustCompile(`\w{4,5}\:\/\/.*?\/`)
	AESIV   = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func GetHost(data string) string {
	allString := HostReg.FindAllString(data, -1)
	return allString[0]
}

func GetReqeust(u string, cookie *map[string]string, header, params map[string]string) (bool, HttpResp) {
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}()

	if u == "" {
		return false, HttpResp{}
	}
	request, _ := http.NewRequest("GET", u, nil)

	if cookie != nil && len(*cookie) != 0 {
		if header == nil {
			header = make(map[string]string)
		}
		header["Cookie"] = CookieMap2Str(*cookie)
	}

	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}
	request.Header.Set(UA_KEY, UA_VALUE)

	//加入get参数
	q := request.URL.Query()
	if params != nil {
		for k, v := range params {
			q.Add(k, v)
		}
	}

	request.URL.RawQuery = q.Encode()
	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//禁止重定向，防止cookie丢失
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	resp, err := client.Do(request)
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil && resp.StatusCode != 301 && resp.StatusCode != 302 {
		return false, HttpResp{Error: err}
	}

	resCookie := resp.Cookies()
	if cookie == nil {
		tempCk := make(map[string]string)
		cookie = &tempCk
	}
	for _, v := range resCookie {
		if _, ok := (*cookie)[v.Name]; ok {
			(*cookie)[v.Name] += ";" + v.Value
		} else {
			(*cookie)[v.Name] = v.Value
		}
	}

	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return false, HttpResp{Error: err2}
	}

	var httpResp HttpResp
	success := false
	//判断是否需要重定向
	localtion, _ := resp.Location()
	if localtion != nil && localtion.String() != "" {
		success, httpResp = GetReqeust(localtion.String(), cookie, nil, nil)
	} else {
		success = true
		httpResp.Error = nil
		httpResp.Data = data
		httpResp.Localtion = resp.Request.URL.String()
		httpResp.Cookie = cookie
	}
	return success, httpResp
}

func GetLocaltion(u string, cookie *map[string]string, header map[string]string) (string, *map[string]string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if u == "" {
		return "", nil
	}
	request, _ := http.NewRequest("GET", u, nil)

	if cookie != nil && len(*cookie) != 0 {
		ck := ""
		if header != nil {
			ck = header["Cookie"]
		} else {
			header = make(map[string]string)
		}
		for k, v := range *cookie {
			ck += k + "=" + v + ";"
		}
		ck = strings.Trim(ck, ";")
		header["Cookie"] = ck
	}

	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}

	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", nil
	}
	_, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", nil
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	resCookie := resp.Cookies()
	for _, v := range resCookie {
		if _, ok := (*cookie)[v.Name]; ok {
			(*cookie)[v.Name] += ";" + v.Value
		} else {
			(*cookie)[v.Name] = v.Value
		}
	}
	return resp.Request.URL.String(), cookie
}

func CookieMap2Str(ck map[string]string) string {
	ckStr := ""
	if ck == nil || len(ck) == 0 {
		return ckStr
	}
	for k, v := range ck {
		ckStr += k + "=" + v + ";"
	}
	return strings.Trim(ckStr, ";")
}

func CookieStr2Map(ck string) map[string]string {
	ckMap := make(map[string]string)
	if ck == "" {
		return nil
	}

	split := strings.Split(ck, ";")
	for _, v := range split {
		split2 := strings.Split(v, "=")
		ckMap[split2[0]] = split2[1]
	}
	return ckMap
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AESEncrypt(text, key []byte) (string, error) {
	//生成cipher.Block 数据块
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("错误 -" + err.Error())
		return "", err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad(text, blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, AESIV)
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func RandStr(n int) string {
	if n < 0 {
		return ""
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func PostForm(api string, cookie *map[string]string, header, params map[string]string) (bool, HttpResp) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	data := make(url.Values)
	if params != nil {
		for k, v := range params {
			data[k] = []string{v}
		}
	}
	request, err := http.NewRequest("POST", api, bytes.NewReader([]byte(data.Encode())))
	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//禁止重定向，防止cookie丢失
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}
	if cookie != nil && len(*cookie) != 0 {
		if header == nil {
			header = make(map[string]string)
		}
		header["Cookie"] = CookieMap2Str(*cookie)
	}

	request.Header.Set(UA_KEY, UA_VALUE)
	request.Header.Set(CONTENT_TYPE_KEY, CAS_CONTENT_TYPE_VALUE)

	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}

	res, err := client.Do(request)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()

	if err != nil && res.StatusCode != 302 && res.StatusCode != 301 {
		return false, HttpResp{}
	}
	respBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Println("读取内容失败:", err.Error())
	}
	resCookie := res.Cookies()
	if cookie == nil {
		tempCk := make(map[string]string)
		cookie = &tempCk
	}
	for _, v := range resCookie {
		if _, ok := (*cookie)[v.Name]; ok {
			(*cookie)[v.Name] += ";" + v.Value
		} else {
			(*cookie)[v.Name] = v.Value
		}
	}

	var resp HttpResp

	location, err := res.Location()
	if location != nil && location.String() != "" {
		resp.Localtion = location.String()
	}
	resp.Cookie = cookie
	resp.Data = respBytes
	resp.Error = nil
	return true, resp
}

func PostJson(api string, cookie *map[string]string, header map[string]string, data interface{}) (bool, HttpResp) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	bytesData := []byte(`{}`)

	if data != nil {
		bytesData, _ = json.Marshal(data)
	}

	request, err := http.NewRequest("POST", api, bytes.NewReader(bytesData))
	timeout := time.Duration(6 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//禁止重定向，防止cookie丢失
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}
	if cookie != nil && len(*cookie) != 0 {
		if header == nil {
			header = make(map[string]string)
		}
		header["Cookie"] = CookieMap2Str(*cookie)
	}

	request.Header.Set(UA_KEY, UA_VALUE)
	request.Header.Set(CONTENT_TYPE_KEY, CAS_CONTENT_TYPE_VALUE)

	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}

	res, err := client.Do(request)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()

	if err != nil && res.StatusCode != 302 && res.StatusCode != 301 {
		return false, HttpResp{}
	}
	respBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, HttpResp{}
	}
	resCookie := res.Cookies()
	if cookie == nil {
		tempCk := make(map[string]string)
		cookie = &tempCk
	}
	for _, v := range resCookie {
		if _, ok := (*cookie)[v.Name]; ok {
			(*cookie)[v.Name] += ";" + v.Value
		} else {
			(*cookie)[v.Name] = v.Value
		}
	}

	var resp HttpResp

	location, err := res.Location()
	if location != nil && location.String() != "" {
		resp.Localtion = location.String()
	}
	resp.Cookie = cookie
	resp.Data = respBytes
	resp.Error = nil
	return true, resp
}

func CheckFileIsExits(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
