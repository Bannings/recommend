package internal_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"recommend/golbal"
	"recommend/log"
	. "recommend/model"
	"time"
)

var (
	operationClient = http.Client{Timeout: time.Duration(300) * time.Millisecond, Transport: t}
)

func GetOperationInfo(info UserInfo, uniqueName string) []Operationinfo {
	conf := golbal.GetConfig()
	url := conf.OperationUrl
	var isGuest string
	if info.Uid > 1000000000 || info.Uid == 0 {
		isGuest = "true"
	} else {
		isGuest = "false"
	}
	var channel string
	if info.Channel == "" {
		channel = "xndm"
	} else {
		channel = info.Channel
	}
	sUrl := url + fmt.Sprintf("?contentSex=%d&isGuest=%s&uid=%d&udid=%s&standUniqueName=%s&channel=%s&appVersion=%s", info.UGender, isGuest, info.Uid, info.UdId, uniqueName, channel, info.AppVersion)
	request, err := http.NewRequest("GET", sUrl, nil)
	if err != nil {
		log.Errorf("get operation requests failed, url:%s", sUrl)
		if uniqueName == "HomeTOPBanner" {
			return getDefaultOperationinfo()
		} else {
			return nil
		}
	}
	var recallResult OperationData
	resp, err := operationClient.Do(request)
	if err != nil {
		if uniqueName == "HomeTOPBanner" {
			return getDefaultOperationinfo()
		} else {
			return nil
		}
	}
	if resp != nil {
		defer resp.Body.Close()
		if 200 == resp.StatusCode {
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil
			}
			err = json.Unmarshal([]byte(respBytes), &recallResult)
			if err != nil {
				log.Errorf("Unmarshal, err:%s", err)
				return nil
			}
		}
	}
	if len(recallResult.Data) > 0 {
		return recallResult.Data
	} else {
		if uniqueName == "HomeTOPBanner" {
			return getDefaultOperationinfo()
		} else {
			return nil
		}
	}
}

func getDefaultOperationinfo() []Operationinfo {
	info1 := Operationinfo{
		OposId:    "4e995c146b6f4066b215aa541a9231f2",
		OposName:  "9680",
		OposOrder: 1,
		ResId:     "0a2d25358ff440708306dbea9a87c646",
		ResourceVO: MgResourceVO{
			ResType: 1,
			URL:     "https://image.samh.xndm.tech/operation/new_test/9680.webp",
			Width:   1200,
			Height:  1100,
		},
		OpId: "ca80a4d22a5c48b6b17cf88e635e1907",
		OperationVO: MgOperationVO{
			OpId:         "ca80a4d22a5c48b6b17cf88e635e1907",
			OpName:       "跳转至阅读浮层",
			OpDesc:       "漫画推荐",
			OpAcitonType: 19,
			OpActionInfo: "tisamanapp://goto?page=readfloat&comic_id=9680",
			EsIds:        "",
		},
		CloseFeatures: 1,
	}
	info2 := Operationinfo{
		OposId:    "f64ff901c3e64c07a96a603de5500587",
		OposName:  "106518",
		OposOrder: 2,
		ResId:     "41d7599382154df1951700ba4b97c538",
		ResourceVO: MgResourceVO{
			ResType: 1,
			URL:     "https://image.samh.xndm.tech/operation/new_test/%E9%80%86%E5%A4%A9%E9%82%AA%E7%A5%9E.webp",
			Width:   1200,
			Height:  1100,
		},
		OpId: "af942f567eb2457f907919f82aea4d3e",
		OperationVO: MgOperationVO{
			OpId:         "af942f567eb2457f907919f82aea4d3e",
			OpName:       "跳转至阅读浮层",
			OpDesc:       "漫画推荐",
			OpAcitonType: 19,
			OpActionInfo: "tisamanapp://goto?page=readfloat&comic_id=106518",
			EsIds:        "",
		},
		CloseFeatures: 1,
	}
	info3 := Operationinfo{
		OposId:    "de14e87ecd71495eb11527f2a186af1f",
		OposName:  "27511",
		OposOrder: 3,
		ResId:     "5a44f6ac0a134db798c52f68652b0e68",
		ResourceVO: MgResourceVO{
			ResType: 1,
			URL:     "http://image.samh.xndm.tech/banner/xbxt.webp",
			Width:   1200,
			Height:  1100,
		},
		OpId: "60cd38eef494423887aac7a15f8dfbf8",
		OperationVO: MgOperationVO{
			OpId:         "60cd38eef494423887aac7a15f8dfbf8",
			OpName:       "跳转至阅读浮层",
			OpDesc:       "漫画推荐",
			OpAcitonType: 19,
			OpActionInfo: "tisamanapp://goto?page=readfloat&comic_id=200000",
			EsIds:        "",
		},
		CloseFeatures: 1,
	}
	info4 := Operationinfo{
		OposId:    "de14e87ecd71495eb11527f2a186af1e",
		OposName:  "27511",
		OposOrder: 3,
		ResId:     "5a44f6ac0a134db798c52f68652b0e69",
		ResourceVO: MgResourceVO{
			ResType: 1,
			URL:     "https://image.samh.xndm.tech/operation/new_test/27511.webp",
			Width:   1200,
			Height:  1100,
		},
		OpId: "60cd38eef494423887aac7a15f8dfbf9",
		OperationVO: MgOperationVO{
			OpId:         "60cd38eef494423887aac7a15f8dfbf9",
			OpName:       "跳转至阅读浮层",
			OpDesc:       "漫画推荐",
			OpAcitonType: 19,
			OpActionInfo: "tisamanapp://goto?page=readfloat&comic_id=27511",
			EsIds:        "",
		},
		CloseFeatures: 1,
	}
	return []Operationinfo{info1, info2, info3, info4}
}
