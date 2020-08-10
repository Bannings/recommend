package middleware

import "github.com/gin-gonic/gin"

func ServerHeader(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
}
