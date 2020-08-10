package middleware

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"recommend/golbal"
	"recommend/log"
	"recommend/model"
	"strconv"
	"strings"
	"sync"
)

var (
	producer     sarama.AsyncProducer
	producerOnce sync.Once
)

func UploadUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isUploadApi(c.Request.URL.Path) {
			c.Next()
			return
		}
		var height, width int
		screenHeight := c.Request.Header.Get("screen_height")
		height, err := strconv.Atoi(screenHeight)
		if err != nil {
			height = 0
		}
		screenWidth := c.Request.Header.Get("screen_width")
		width, err = strconv.Atoi(screenWidth)
		if err != nil {
			width = 0
		}
		suid := c.Query("uid")
		var uid int64
		uid, err = strconv.ParseInt(suid, 10, 64)
		if err != nil {
			uid = 0
		}
		gender := c.Query("gender")
		var genderId int
		switch gender {
		case "0":
			genderId = 1
		case "1":
			genderId = 2
		default:
			genderId = 0
		}
		info := model.UserBehaviorInfo{
			Access:       c.Request.Header.Get("access"),
			AppVersion:   c.Request.Header.Get("app_version"),
			Channel:      c.Request.Header.Get("channel"),
			DeviceId:     c.Request.Header.Get("device_id"),
			OS:           c.Request.Header.Get("os"),
			OSVersion:    c.Request.Header.Get("os_version"),
			Carrier:      c.Request.Header.Get("carrier"),
			ScreenHeight: height,
			ScreenWidth:  width,
			DeviceModel:  c.Request.Header.Get("device_model"),
			DeviceBrand:  c.Request.Header.Get("device_brand"),
			Gender:       genderId,
			IP:           getClientIP(c),
			Uid:          uid,
		}
		codeInfo, _ := json.Marshal(info)
		message := &sarama.ProducerMessage{Topic: "user_request_upload", Key: sarama.StringEncoder("info"), Value: sarama.StringEncoder(string(codeInfo))}
		KafkaConnection().Input() <- message
	}
}

func isUploadApi(path string) bool {
	if path == "/top/homepage" {
		return true
	}
	if path == "/top/v2/homepage" {
		return true
	}
	if path == "/top/v3/recommend_stream" {
		return true
	}
	return false
}

func getClientIP(ctx *gin.Context) string {

	ip := ctx.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = ctx.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	ips := strings.Split(ip, ":")
	if len(ips) > 0 {
		ip = ips[0]
	}
	if ip == "" {
		return "127.0.0.1"
	}

	return ip
}

func KafkaConnection() sarama.AsyncProducer {
	producerOnce.Do(func() {
		config := sarama.NewConfig()
		config.Producer.Timeout = 10
		var err error
		conf := golbal.GetConfig()
		producer, err = sarama.NewAsyncProducer(conf.Kafka, config)
		if err != nil {
			log.Error(err)
			//errors_.CheckErrSendEmail(err)
		}
	})
	return producer
}

func HandleError() {
	for {
		select {
		case err := <-KafkaConnection().Errors():
			log.Error(err)
			//errors_.CheckErrSendEmail(err)
			continue
		}

	}

}
