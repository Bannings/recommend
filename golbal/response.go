package golbal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Success        = &Response{Status: "0", Message: "ok", CloseMenu: 1}
	InvalidRequest = &Response{Status: "400", Message: "无效的请求."}
	InternalError  = &Response{Status: "500", Message: "服务器内部错误."}
	DataDeplete    = &Response{Status: "2001", Message: "无更多数据"}
)

type Response struct {
	Status    string      `json:"status"`
	Message   string      `json:"msg"`
	Data      interface{} `json:"data"`
	CloseMenu int         `json:"close_menu"`
}

func SendResponse(c *gin.Context, response Response, data interface{}) {
	if data != nil {
		response.Data = data
	} else {
		response.Data = errorData{response.Message}
	}
	c.JSON(http.StatusOK, response)
}

type errorData struct {
	info string
}
