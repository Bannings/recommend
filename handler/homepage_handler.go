package handler

import (
	"github.com/gin-gonic/gin"
	"recommend/golbal"
	"recommend/internal_api"
	"recommend/log"
	. "recommend/model"
	"recommend/util"
	"recommend/util/timed_task"
	"strconv"
	"strings"
)

func HomepageTopRecommend(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}

	userInfo := getUserInfo(c)
	userInfo.UserDeviceInfo = userDeviceInfo
	HomepageRequestInfoLog(c, userInfo.Uid, userInfo.UdId, userInfo.Gender)
	var result []SectionDataConfig
	result = util.HomepageTopApplyV2(userInfo)
	golbal.SendResponse(c, *golbal.Success, result)
	return

}

func HomepageTopRecommendV2(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}

	userInfo := getUserInfo(c)
	userInfo.UserDeviceInfo = userDeviceInfo
	HomepageRequestInfoLog(c, userInfo.Uid, userInfo.UdId, userInfo.Gender)
	var result []SectionDataConfig
	result = util.HomepageTopApplyV4(userInfo)
	golbal.SendResponse(c, *golbal.Success, result)
	return

}

func HomepageTopRecommendV3(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}

	userInfo := getUserInfo(c)
	userInfo.UserDeviceInfo = userDeviceInfo
	HomepageRequestInfoLog(c, userInfo.Uid, userInfo.UdId, userInfo.Gender)
	var result []SectionDataConfig
	result = util.HomepageTopApplyV4(userInfo)
	golbal.SendResponse(c, *golbal.Success, result)
	return

}

func HomepageTopMoreV2(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	userInfo.UserDeviceInfo = userDeviceInfo
	sectionId := c.DefaultQuery("section_id", "1")
	num := c.Query("page_num")
	size := c.Query("page_size")
	pageNum, err := strconv.Atoi(num)
	if err != nil {
		pageNum = 0
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		pageSize = 18
	}
	jsonData := util.HomepageMore(userInfo, sectionId, pageNum, pageSize)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func HomepageTopRenew(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	userInfo.UserDeviceInfo = userDeviceInfo
	sectionId := c.DefaultQuery("section_id", "1")
	num := c.Query("page_num")
	size := c.Query("page_size")
	pageNum, err := strconv.Atoi(num)
	if err != nil {
		pageNum = 0
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		pageSize = 18
	}
	jsonData := util.HomepageRenew(userInfo, sectionId, pageNum, pageSize)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetBillBoard(c *gin.Context) {
	golbal.SendResponse(c, *golbal.Success, nil)
	return
}

func GetBillBoardV2(c *gin.Context) {
	golbal.SendResponse(c, *golbal.Success, nil)
	return
}

func GetRelateRecommend(c *gin.Context) {
	var recommendReq RelateRecommendReq
	if err := c.Bind(&recommendReq); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, "Missing cid")
		return
	}
	var sGender string
	if recommendReq.Gender == 0 {
		sGender = golbal.MALE_STR
	} else {
		sGender = golbal.FEMALE_STR
	}
	info := UserInfo{Uid: recommendReq.Uid, Gender: recommendReq.Gender, SGender: sGender, Platform: "samh", Expose: recommendReq.Expose}
	var jsonData *SectionDataConfig
	if recommendReq.ReadingCount >= 5 {
		jsonData = util.GetRelateRecommendSection(info, recommendReq.ComicName, recommendReq.Cid, recommendReq.SectionOrder)
	} else {
		jsonData = util.GetInterestRecommendSection(info, recommendReq.ComicName, recommendReq.Cid, recommendReq.SectionOrder)
	}

	if jsonData != nil {
		golbal.SendResponse(c, *golbal.Success, jsonData)
		return
	} else {
		data := SectionDataConfig{}
		golbal.SendResponse(c, *golbal.Success, data)
		return
	}
}

func getUserInfo(c *gin.Context) UserInfo {
	uid := c.Query("uid")
	uidInt64, _ := strconv.ParseInt(uid, 10, 0)
	gender := c.Query("gender")
	udid := c.Query("udid")
	deviceBrand := c.Request.Header.Get("device_brand")
	if udid == "" {
		udid = c.Query("device_id")
	}

	var sGender, os string
	var genderId int
	if gender == "0" {
		sGender = golbal.MALE_STR
		genderId = golbal.MALE_ID
	} else {
		sGender = golbal.FEMALE_STR
		genderId = golbal.FEMALE_ID
	}
	platform := getPlatform(c)
	if platform == golbal.BANNER_PLATFORM_ANDROID {
		os = golbal.BANNER_PLATFORM_ANDROID_STR
	} else if platform == golbal.BANNER_PLATFORM_IOS {
		os = golbal.BANNER_PLATFORM_IOS_STR
	}
	sessionId := c.Query("sessionid")
	version := c.Query("version")
	if uidInt64 == 0 {
		return UserInfo{Uid: 0, IsNewer: true, Gender: genderId, SGender: sGender, DeviceType: platform, OS: os, DeviceBrand: deviceBrand, Platform: "samh", AppVersion: version}
	}
	userInfoResponse, err := internal_api.GetUserPortrait(uidInt64)
	if err != nil {
		log.Error(err)
		userInfoResponse = &internal_api.RecAppJson{}
	}
	info := UserInfo{
		Uid:          uidInt64,
		UdId:         udid,
		Gender:       genderId,
		UGender:      genderId,
		DeviceBrand:  deviceBrand,
		OS:           os,
		Platform:     "samh",
		AppVersion:   version,
		SessionId:    sessionId,
		ReadCartoons: userInfoResponse.Data.UAction.URead,
		SGender:      sGender,
		IsNewer:      userInfoResponse.Data.USta.UIsNewer,
		Age:          userInfoResponse.Data.USta.UAge,
		DeviceType:   platform,
	}
	return info
}

func getPlatform(c *gin.Context) (platform int) {
	platformStr := strings.ToLower(c.Query("platform"))
	if platformStr == "" {
		platformStr = strings.ToLower(c.Query("platformname"))
	}
	if platformStr != "" {
		switch platformStr {
		case "android":
			platform = golbal.BANNER_PLATFORM_ANDROID
		case "ios":
			platform = golbal.BANNER_PLATFORM_IOS
		case "wxapp":
			platform = golbal.BANNER_PLATFORM_WXAPP
		}
	}
	return platform
}

func headerDefaultGet(c *gin.Context, header, defaultVal string) string {
	headerVal := c.Request.Header.Get(header)
	if headerVal == "" {
		headerVal = defaultVal
	}
	return headerVal
}

func HomepageRequestInfoLog(c *gin.Context, uid int64, udid string, gender int) {
	platformName := c.DefaultQuery("platformname", "-")
	productName := c.DefaultQuery("productname", "-")
	userAgent := c.Request.UserAgent()
	if userAgent == "" {
		userAgent = "-"
	}
	androidId := c.DefaultQuery("android_id", "-")
	appVersion := headerDefaultGet(c, "app_version", "-")
	deviceName := c.DefaultQuery("device_model", "-")
	oaid := c.DefaultQuery("oaid", "-")
	systemVersion := c.DefaultQuery("os_version", "-")
	channel := c.DefaultQuery("client-channel", "-")
	access := c.DefaultQuery("access", "-")
	log.DataLogger.Infof(`uid=%d udid=%s gender=%d platformname=%s productname=%s useragent="%s" android_id=%s app_version=%s device_name=%s system_version=%s oaid=%s channel=%s access=%s`,
		uid, udid, gender, platformName, productName, userAgent, androidId, appVersion, deviceName, systemVersion, oaid, channel, access)
}

func BannerReload(c *gin.Context) {
	timed_task.ReloadBannerData()
	golbal.SendResponse(c, *golbal.Success, nil)
	return
}

func GetBookList(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	pageNum := c.Query("page_num")
	userInfo.UserDeviceInfo = userDeviceInfo
	userInfo.PageNum = pageNum
	result, flag := util.GetBookList(userInfo)
	if flag == 0 {
		golbal.SendResponse(c, *golbal.DataDeplete, result)
	} else {
		golbal.SendResponse(c, *golbal.Success, result)
	}
	return
}

func GetBookListV2(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	pageNum := c.Query("page_num")
	userInfo.UserDeviceInfo = userDeviceInfo
	userInfo.PageNum = pageNum
	result, flag := util.GetBookListV2(userInfo)
	if flag == 0 {
		golbal.SendResponse(c, *golbal.DataDeplete, result)
	} else {
		golbal.SendResponse(c, *golbal.Success, result)
	}
	return
}
