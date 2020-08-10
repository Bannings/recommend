package internal_api

import (
	"recommend/model"
	"testing"
)

func TestGetCoverRPC(t *testing.T) {
	info := model.UserInfo{Uid: 89567549, UdId: "65466789321", Gender: 0}
	a := GetStreamRecommend(info)
	t.Log(a)
}
