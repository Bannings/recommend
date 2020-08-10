package handler

import (
	"github.com/gin-gonic/gin"
	"recommend/golbal"
)

//已废弃api
func NewUserInterestCtypesRecordApi(c *gin.Context) {
	golbal.SendResponse(c, *golbal.Success, nil)
	return
}

func NewUserInterestCtypesApi(c *gin.Context) {
	golbal.SendResponse(c, *golbal.Success, golbal.NewUserInterestCtypes)
	return
}
