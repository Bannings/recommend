package internal_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"recommend/golbal"
	"recommend/log"
	"recommend/model"
	"strconv"
	"time"
)

var (
	t = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   400,
		MaxIdleConns:          400,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	}

	userPortraitClient = http.Client{Timeout: time.Duration(1) * time.Second, Transport: t}
)

type UserAction struct {
	UClick   string `json:"ucid_click"`   //   25934|7119
	URead    string `json:"ucid_read"`    //   25934|7119
	USearch  string `json:"ucid_search"`  //   25934|7119
	UCollect string `json:"ucid_collect"` //   25934|7119
	UPay     string `json:"ucid_pay"`     //   25934|7119
	UInter   string `json:"uinterest"`    //   xuanhuan,玄幻|rexue,热血
}

type UserSta struct {
	// 基础信息
	UGender  string `json:"ugender"` // 男:1 女:2 未知:0
	UAge     string `json:"uage"`    // 未成年:minor 成年人:adult
	UIsNewer bool   `json:"unewer"`  // 新用户：true 老用户：false

	// 设备信息
	UOs           string `json:"uos"` //android ios
	UOsVersion    string `json:"uos_version"`
	UScreenWidth  string `json:"uscreen_width"`
	UScreenheight string `json:"uscreen_height"`
	UChannel      string `json:"uchannel"`
	UDeviceBrand  string `json:"udevice_brand"`
	UDeviceModel  string `json:"udevice_model"`
	UAccess       string `json:"uaccess"`
	UAppVersion   string `json:"uapp_version"`

	// 地域信息
	UCapital string `json:"ucapital"` // 返回中文，英文容易有歧义：广州
	UCity    string `json:"ucity"`    // 返回中文，英文容易有歧义：广州

	// 登录信息
	UBind string `json:"ubind_info"` // 用户登录方式

	// 唤醒信息
	UFirstTime string `json:"ufirst_time"` // 时间戳
	ULastTime  string `json:"ulast_time"`  // 时间戳

	// vip信息
	UIsVip       string `json:"uis_vip"`
	UVipLevel    string `json:"uvip_level"`
	UPayPlatform string `json:"upay_platform"` // alipay  applepay  huaweipay  weixinpay qqpay
}

type RedisUserInfoResponse struct {
	Uid     int64      `json:"uid"`
	UAction UserAction `json:"user_action"`
	USta    UserSta    `json:"user_statics"`
}

type RecAppJson struct {
	Status int                   `json:"status"`
	Msg    string                `json:"msg"`
	Data   RedisUserInfoResponse `json:"data"`
}

type Request struct {
	Uid          string `url:"uid"`
	Recall       string `url:"recall"`
	DeviceId     string `url:"device_id"`
	UdId         string `url:"udid"`
	DeviceBrand  string `url:"device_brand"`
	DeviceModel  string `url:"device_model"`
	ScreenWidth  string `url:"screen_width"`
	ScreenHeight string `url:"screen_height"`
	Carrier      string `url:"carrier"`
	Access       string `url:"access"`
	Os           string `url:"os"`
	OsVersion    string `url:"os_version"`
	Channel      string `url:"channel"`
	AppVersion   string `url:"app_version"`
	UAddr        string `url:"addr"`
	Gender       int    `url:"gender"`
	Sgender      string `url:"sgender"`
	Algo         string `url:"algo"`
	SectionID    int    `url:"section_id"`
	Len          int    `url:"len"`
}

func (r *Request) fromUserInfo(userInfo model.UserInfo) {
	r.Uid = strconv.FormatInt(userInfo.Uid, 10)
	r.Sgender = strconv.Itoa(userInfo.Gender)
	r.Gender = userInfo.Gender
	r.Recall = userInfo.Recall
	r.UdId = userInfo.UdId
	r.DeviceId = userInfo.UdId
	r.DeviceBrand = userInfo.DeviceBrand
	r.DeviceModel = userInfo.DeviceModel
	r.ScreenWidth = userInfo.ScreenWidth
	r.ScreenHeight = userInfo.ScreenHeight
	r.Carrier = userInfo.Carrier
	r.Access = userInfo.Access
	r.Os = userInfo.Os
	r.OsVersion = userInfo.OsVersion
	r.Channel = userInfo.Channel
	r.AppVersion = userInfo.AppVersion
	r.UAddr = userInfo.UAddr
	r.Len = userInfo.Len
	r.Algo = userInfo.Algo
}

func Struct2Url(stru interface{}, baseUrl string) string {

	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Error(err)
	}
	q := u.Query()

	value := reflect.ValueOf(stru)
	typ := reflect.TypeOf(stru)
	for i := 0; i < typ.NumField(); i++ {
		var name string
		name, ok := typ.Field(i).Tag.Lookup("url")
		if !ok || name == "-" {
			continue
		}
		var fieldVal string
		switch typV := value.Field(i).Interface().(type) {
		case string:
			fieldVal = typV
		case int:
			fieldVal = strconv.Itoa(typV)
		}
		q.Add(name, fieldVal)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// GetUserInfo 用于获取用户信息并存入结构体
func GetUserInfo(userId int64, udId, deviceBrand string, gender, platform int, userDeviceInfo model.UserDeviceInfo) (userInfo *model.UserInfo) {
	userInfoResponse, err := GetUserPortrait(userId)
	if err != nil {
		log.Error(err)
		//errors_.CheckCommonErr(err)
		return nil
	}

	userInfo.UdId = udId
	userInfo.Uid = userId
	userInfo.Gender = gender
	userInfo.DeviceBrand = deviceBrand
	userInfo.Platform = "samh"
	userInfo.ReadCartoons = userInfoResponse.Data.UAction.URead
	if gender == golbal.FEMALE_ID {
		userInfo.SGender = golbal.FEMALE_STR
	} else {
		userInfo.SGender = golbal.MALE_STR
	}
	if platform == golbal.BANNER_PLATFORM_ANDROID {
		userInfo.OS = golbal.BANNER_PLATFORM_ANDROID_STR
	} else if platform == golbal.BANNER_PLATFORM_IOS {
		userInfo.OS = golbal.BANNER_PLATFORM_IOS_STR
	}
	if err == nil {
		userInfo.IsNewer = userInfoResponse.Data.USta.UIsNewer
		userInfo.Age = userInfoResponse.Data.USta.UAge
	}
	userInfo.UserDeviceInfo = userDeviceInfo
	return userInfo
}

func GetUserPortrait(uid int64) (*RecAppJson, error) {
	conf := golbal.GetConfig()
	urls := fmt.Sprintf(conf.UserPortraitURL+"?uid=%d&platform=samh", uid)
	request, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return nil, err
	}
	var userInfoResponse RecAppJson

	resp, err := userPortraitClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
		if 200 == resp.StatusCode {
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(respBytes), &userInfoResponse)
			if err != nil {
				return nil, err
			}
		}
	}
	return &userInfoResponse, err
}
