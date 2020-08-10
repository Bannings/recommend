package handler

import (
	"github.com/gin-gonic/gin"
	"recommend/golbal"
	. "recommend/model"
	"recommend/util"
)

func HomepageRecommendStreamWrapper(c *gin.Context) {

	userInfo := getUserInfo(c)
	pageSize := 18
	jsonData := util.SamhRecommendStreamApplyV2(userInfo, pageSize)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func HomepageRecommendStreamV3Wrapper(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 18
	}
	pageNum := request.PageNum
	if pageNum == 0 {
		pageNum = 1
	}
	jsonData := util.SamhRecommendStreamApplyV3(userInfo, pageSize, pageNum)
	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFromTypeRecommendDefault(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}

	userInfo := getUserInfo(c)
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = 18
	}
	jsonData := util.SamhTypeApply(userInfo, pageSize, golbal.INDEX_TYPE)

	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFeedStream(c *gin.Context) {
	var request UserRequest
	if err := c.Bind(&request); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, nil)
		return
	}
	userInfo := getUserInfo(c)
	var jsonData []SectionDataConfig
	jsonData = util.GetFeedStream(userInfo, request.PageNum)
	//if request.SessionId == "" {
	//	jsonData = util.GetFeedStream(userInfo, request.PageNum)
	//} else {
	//	jsonData = util.GetInternalFeedStream(userInfo.Uid, request.SessionId, userInfo.Gender, request.PageNum)
	//}

	golbal.SendResponse(c, *golbal.Success, jsonData)
	return
}

func GetFeedStreamRelateRecommend(c *gin.Context) {
	var recommendReq RelateRecommendReq
	if err := c.Bind(&recommendReq); err != nil {
		golbal.SendResponse(c, *golbal.InvalidRequest, "Missing cid")
		return
	}
	data := SectionDataConfig{}
	golbal.SendResponse(c, *golbal.Success, data)
	return
	//var sGender string
	//var jsonData *SectionDataConfig
	//if recommendReq.Gender == 0 {
	//	sGender = gobal.MALE_STR
	//} else {
	//	sGender = gobal.FEMALE_STR
	//}
	//info := UserInfo{Uid: recommendReq.Uid, Gender: recommendReq.Gender, SGender: sGender, Platform: "samh", Expose: recommendReq.Expose}
	//if recommendReq.SessionId == "" {
	//	jsonData = util.GetFeedStreamRelate(info, recommendReq.Cid)
	//} else {
	//	jsonData = util.GetInternalFeedStreamRelate(recommendReq.Uid, recommendReq.Cid, recommendReq.SessionId, recommendReq.Gender)
	//}
	//
	//if jsonData != nil {
	//	gobal.SendResponse(c, *gobal.Success, jsonData)
	//	return
	//} else {
	//	data := SectionDataConfig{}
	//	gobal.SendResponse(c, *gobal.Success, data)
	//	return
	//}
}
