package main

type SchoolList struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    []struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		TenantCode  string  `json:"tenantCode"`
		Img         string  `json:"img"`
		Distance    string  `json:"distance"`
		Dis         float64 `json:"dis"`
		IdsURL      string  `json:"idsUrl"`
		JoinType    string  `json:"joinType"`
		AppID       string  `json:"appId"`
		CasLoginURL string  `json:"casLoginUrl"`
		IsEnter     int     `json:"isEnter"`
	} `json:"data"`
}

type SchoolInfo struct {
	ErrCode int              `json:"errCode"`
	ErrMsg  string           `json:"errMsg"`
	Data    []SchoolInfoData `json:"data"`
}

type SchoolInfoData struct {
	ID                         string        `json:"id"`
	Name                       string        `json:"name"`
	TenantCode                 string        `json:"tenantCode"`
	Img                        string        `json:"img"`
	Distance                   string        `json:"distance"`
	ShortName                  string        `json:"shortName"`
	JoinType                   string        `json:"joinType"`
	CasLoginURL                string        `json:"casLoginUrl"`
	IsEnter                    int           `json:"isEnter"`
	IdsURL                     string        `json:"idsUrl"`
	AmpURL                     string        `json:"ampUrl"`
	AmpURL2                    string        `json:"ampUrl2"`
	PriorityURL                string        `json:"priorityUrl"`
	AppID                      string        `json:"appId"`
	AppSecret                  string        `json:"appSecret"`
	MsgURL                     string        `json:"msgUrl"`
	MsgAccessToken             string        `json:"msgAccessToken"`
	MsgAppID                   string        `json:"msgAppId"`
	MsgAppIDIos                string        `json:"msgAppIdIos"`
	ZgAppKey                   string        `json:"zgAppKey"`
	YktBalanceURL              string        `json:"yktBalanceUrl"`
	YktTransferURL             string        `json:"yktTransferUrl"`
	YktQrCodeURL               string        `json:"yktQrCodeUrl"`
	XykURL                     string        `json:"xykUrl"`
	UserShowCollege            string        `json:"userShowCollege"`
	ScheduleOpenURL            string        `json:"scheduleOpenUrl"`
	ScheduleDataURL            string        `json:"scheduleDataUrl"`
	IsIdsProxy                 string        `json:"isIdsProxy"`
	TenantNameImg              string        `json:"tenantNameImg"`
	IsNeedAlias                string        `json:"isNeedAlias"`
	ModifyPassURL              string        `json:"modifyPassUrl"`
	ModifyPassSuccessURL       string        `json:"modifyPassSuccessUrl"`
	ModifyPassDescr            string        `json:"modifyPassDescr"`
	TaskURL                    string        `json:"taskUrl"`
	TaskAppID                  string        `json:"taskAppId"`
	CircleShowType             string        `json:"circleShowType"`
	IsShowHotList              string        `json:"isShowHotList"`
	AppStyleVersionID          string        `json:"appStyleVersionId"`
	AppStyleResURL             string        `json:"appStyleResUrl"`
	LikeBtnSpace               string        `json:"likeBtnSpace"`
	IRobotURL                  string        `json:"iRobotUrl"`
	ServicePagePlace           string        `json:"servicePagePlace"`
	ScheduleAllDataURL         string        `json:"scheduleAllDataUrl"`
	ScheduleUpdateDataURL      string        `json:"scheduleUpdateDataUrl"`
	ShopURL                    string        `json:"shopUrl"`
	HomePageDisplayItem        string        `json:"homePageDisplayItem"`
	TaoBannerID                string        `json:"taoBannerId"`
	CanIdsLogin                string        `json:"canIdsLogin"`
	AppCacheDisable            string        `json:"appCacheDisable"`
	HomePageDisplayItemTeacher string        `json:"homePageDisplayItemTeacher"`
	LossPwdDesc                string        `json:"lossPwdDesc"`
	IsAmpProxy                 string        `json:"isAmpProxy"`
	ProvinceID                 string        `json:"provinceId"`
	YbSwitch                   string        `json:"ybSwitch"`
	Amp3URL                    string        `json:"amp3Url"`
	IsOpenFission              string        `json:"isOpenFission"`
	IsOpenOauth                string        `json:"isOpenOauth"`
	AmpRobotURL                string        `json:"ampRobotUrl"`
	MediaVersion               string        `json:"mediaVersion"`
	BadHTTPSBlock              string        `json:"badHttpsBlock"`
	FaqForumID                 string        `json:"faqForumId"`
	CampusReqProxy             string        `json:"campusReqProxy"`
	AppStoreURL                string        `json:"appStoreUrl"`
	StudentVersion             string        `json:"studentVersion"`
	CircleCanSeeOffCampus      string        `json:"circleCanSeeOffCampus"`
	ContactDisplayItem         string        `json:"contactDisplayItem"`
	ContactDisplayItemTeacher  string        `json:"contactDisplayItemTeacher"`
	HomeFirstShow              string        `json:"homeFirstShow"`
	AllowSendMsg               int           `json:"allowSendMsg"`
	TeacherVersion             string        `json:"teacherVersion"`
	YibanBuild                 int           `json:"yibanBuild"`
	FreshPostRange             string        `json:"freshPostRange"`
	YibanAuthType              string        `json:"yibanAuthType"`
	CanInteractive             int           `json:"canInteractive"`
	SecondHandSwitch           string        `json:"secondHandSwitch"`
	YwtStatus                  string        `json:"ywtStatus"`
	YwtPrefixURL               string        `json:"ywtPrefixUrl"`
	YwtServiceURL              string        `json:"ywtServiceUrl"`
	CollegeTown                string        `json:"collegeTown"`
	HasOpenMessageFresh        string        `json:"hasOpenMessageFresh"`
	TeacherHomePages           []interface{} `json:"teacherHomePages"`
	StudentHomePages           []int         `json:"studentHomePages"`
	NoAuthHomePages            []int         `json:"noAuthHomePages"`
}

type HttpResp struct {
	Data      []byte
	Cookie    *map[string]string
	Localtion string
	Error     error
}
