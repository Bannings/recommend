package handler

import (
	"github.com/gin-gonic/gin"
	"recommend/golbal"
	. "recommend/model"
	"recommend/util"
)

//成都迁移之前接口，仅2.0.17以下版本使用

func GetFromPersonalRecommend(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	jsonData := util.SamhApply(userInfo, golbal.INDEX_TYPE)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFromVipRecommend(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	jsonData := util.SamhApply(userInfo, golbal.INDEX_VIP)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFromNewRecommend(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	jsonData := util.SamhApply(userInfo, golbal.INDEX_NEW)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFromUniqueRecommend(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	jsonData := util.SamhApply(userInfo, golbal.INDEX_UNIQUE)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFromQualityRecommend(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	jsonData := util.SamhApply(userInfo, golbal.INDEX_QUALITY)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFromHighlyRecommend(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	jsonData := util.SamhApply(userInfo, golbal.INDEX_HIGHLY)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetHomepageTopRecommend(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}

	userInfo := getUserInfo(c)
	userInfo.UserDeviceInfo = userDeviceInfo

	HomepageRequestInfoLog(c, userInfo.Uid, userInfo.UdId, userInfo.Gender)
	jsonData := util.HomepageTopApplyV2(userInfo)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}
