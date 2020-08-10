package model

type UserBehaviorInfo struct {
	Uid          int64  `json:"uid"`
	Gender       int    `json:"selected_gender"`
	DeviceId     string `json:"device_id"`
	DeviceBrand  string `json:"device_brand"`
	DeviceModel  string `json:"device_model"`
	ScreenWidth  int    `json:"screen_width"`
	ScreenHeight int    `json:"screen_height"`
	Carrier      string `json:"carrier"`
	Access       string `json:"access"`
	OS           string `json:"os"`
	OSVersion    string `json:"os_version"`
	Channel      string `json:"channel"`
	AppVersion   string `json:"app_version"`
	IP           string `json:"ip"`
}

type UserRequest struct {
	Uid       int64  `json:"uid" form:"uid"`
	Gender    int    `json:"gender" form:"gender"`
	DeviceId  string `json:"device_id" form:"device_id"`
	PageSize  int    `json:"page_size" form:"page_size"`
	PageNum   int    `json:"page_num" form:"page_num"`
	Udid      string `json:"udid" form:"udid"`
	SessionId string `json:"sessionid" form:"sessionid"`
}
