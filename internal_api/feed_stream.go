package internal_api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"recommend/golbal"
	"recommend/log"
	. "recommend/model"
	"time"
)

var (
	streamClient = http.Client{Timeout: time.Duration(300) * time.Millisecond, Transport: t}
)

type StreamReq struct {
	Uid     int64  `json:"uid"`
	Udid    string `json:"udid"`
	Gender  int    `json:"gender"`
	TakeLen int    `json:"take_len"`
}

type StreamResponsestruct struct {
	Msg  string     `json:"msg"`
	Flag int        `json:"flag"`
	Data StreamData `json:"data"`
}

type StreamData struct {
	Code        int      `json:"code"`
	Result      []string `json:"result"`
	Passthrough ABExps   `json:"passthrough"`
}

func GetStreamRecommend(info UserInfo) []string {
	conf := golbal.GetConfig()
	url := conf.FeedStreamUrl
	//url := "http://test.samh.xndm.tech/feeds_recommend/v1"
	req := StreamReq{Uid: info.Uid, Udid: info.UdId, Gender: info.Gender, TakeLen: 8}
	bytesData, err := json.Marshal(req)
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Errorf("get stream recommend failed, url:%s", url)
		return nil
	}
	var StreamResult StreamResponsestruct
	resp, err := streamClient.Do(request)
	if err != nil {
		log.Errorf("get stream recommend failed, err:%s", err)
		return nil
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(respBytes, &StreamResult)
	if len(StreamResult.Data.Result) > 0 {
		return StreamResult.Data.Result
	} else {
		log.Error("get stream recommend failed, return nil")
		return nil
	}
}
