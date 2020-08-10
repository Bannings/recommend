package model

type UserDeviceInfo struct {
	DeviceBrand  string `json:"device_brand" url:"device_brand" form:"device_brand"`
	DeviceModel  string `json:"device_model" url:"device_model" form:"device_model"`
	ScreenWidth  string `json:"screen_width" url:"screen_width" form:"screen_width"`
	ScreenHeight string `json:"screen_height" url:"screen_height" form:"screen_height"`
	Carrier      string `json:"carrier" url:"carrier" form:"carrier"`
	Access       string `json:"access" url:"access" form:"access"`
	Os           string `json:"os" url:"os" form:"os"`
	OsVersion    string `json:"os_version" url:"os_version" form:"os_version"`
	Channel      string `json:"channel" url:"channel" form:"channel"`
	AppVersion   string `json:"version" url:"version" form:"version"`
	UAddr        string `json:"uaddr" url:"uaddr"`
}

type UserInfo struct {
	Uid          int64    `url:"uid"`
	UdId         string   `url:"udid"`
	Gender       int      `url:"gender"`
	UGender      int      `url:"ugender"`
	OS           string   `url:"os"`
	Age          string   `url:"-"`
	IsNewer      bool     `url:"-"`
	Recall       string   `url:"recall"`
	ReadCartoons string   `url:"_"`
	Interest     string   `url:"-"`
	Algo         string   `url:"algo"`
	DeviceBrand  string   `url:"device_brand"`
	Len          int      `url:"len"`
	SGender      string   `url:"sgender"`
	Platform     string   `url:"platform"`
	SessionId    string   `url:"_"`
	DeviceType   int      `url:"_"` //BANNER_PLATFORM_ANDROID|BANNER_PLATFORM_IOS|BANNER_PLATFORM_WXAPP
	AppVersion   string   `url:"_"`
	PageNum      string   `url:"_"`
	Expose       []string `url:"_"`
	UserDeviceInfo
}

type NewUserIDReceive struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		NewerID string `json:"newer_id"`
		GuestID string `json:"guest_id"`
	} `json:"data"`
}

type NewUserInterest struct {
	Uid    int64  `form:"uid" json:"uid" binding:"required"`
	Ctypes string `form:"comic_types" json:"comic_types" binding:"required"`
}

type NewUserID struct {
	NewerID int64
	GuestID int64
}

type ReadHistory struct {
	Uid     int64  `json:"luid" gorm:"column:LUid"`
	Cid     int64  `json:"lcid" gorm:"column:LCid"`
	Chapter int64  `json:"lchapter" gorm:"column:Lchapter"`
	Time    string `json:"ltime" gorm:"column:Ltime"`
}

func (history *ReadHistory) TableName() string {
	return "read_log_main"
}

type HistoryDeatil struct {
	Chapter     string
	ChapterName string
}
