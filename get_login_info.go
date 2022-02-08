package main

import (
	"encoding/json"
	"strings"
)

func GetSchoolLoginInfo(schoolName string) map[string]string {
	id, ck := GetSchoolId(schoolName)
	if id == "" {
		return nil
	}
	info, _ := GetSchoolInfo(id, ck)
	if info == nil {
		return nil
	}
	clientUrl := ""
	if strings.Contains(info.AmpURL, "campusphere") {
		clientUrl = info.AmpURL
	} else if strings.Contains(info.AmpURL2, "campusphere") {
		clientUrl = info.AmpURL2
	}
	header := make(map[string]string)
	header[UA_KEY] = UA_VALUE
	location, newCookie := GetLocaltion(clientUrl, ck, header)
	if newCookie == nil {
		return nil
	}
	result := make(map[string]string)
	result["campusHost"] = GetHost(clientUrl)
	result["loginUrl"] = location
	result["cookie"] = CookieMap2Str(*newCookie)
	result["loginHost"] = GetHost(location)
	result["loginType"] = info.JoinType
	return result
}

func GetSchoolId(schoolName string) (string, *map[string]string) {
	success, resp := GetReqeust(SCHOOL_LIST_API, nil, nil, nil)
	if !success {
		return "", nil
	}
	var schoolList SchoolList
	err := json.Unmarshal(resp.Data, &schoolList)
	if err != nil {
		return "", nil
	}
	if schoolList.ErrCode != 0 {
		return "", nil
	}

	for _, v := range schoolList.Data {
		if v.Name == schoolName {
			return v.ID, resp.Cookie
		}
	}
	return "", nil
}

func GetSchoolInfo(schoolId string, cookie *map[string]string) (*SchoolInfoData, *map[string]string) {
	params := make(map[string]string)
	params["ids"] = schoolId
	success, resp := GetReqeust(SCHOOL_INFO_API, cookie, nil, params)
	if !success {
		return nil, nil
	}
	var schoolInfo SchoolInfo
	err := json.Unmarshal(resp.Data, &schoolInfo)
	if err != nil {
		return nil, nil
	}

	if schoolInfo.ErrCode != 0 {
		return nil, nil
	}
	return &schoolInfo.Data[0], resp.Cookie
}
