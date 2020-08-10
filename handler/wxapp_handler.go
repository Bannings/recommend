package handler

import (
	"github.com/gin-gonic/gin"
	"recommend/golbal"
	. "recommend/model"
	"recommend/util"
	"strconv"
)

func GetHomepageTopRecommendV2Wxapp(c *gin.Context) {

	userInfo := getUserInfo(c)
	jsonData := util.HomepageTopWxappApplyV2(userInfo)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetHomepageTopMoreV2Default(c *gin.Context) {
	var userDeviceInfo UserDeviceInfo
	if err := c.Bind(&userDeviceInfo); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	var abTestInd string
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
	jsonData := util.WxappHomepageMoreApplyV2(userInfo, sectionId, abTestInd, pageNum, pageSize)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}
