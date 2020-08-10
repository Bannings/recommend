package middleware

import (
	"github.com/gin-gonic/gin"
	"recommend/log"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		referer := c.Request.Referer()
		if referer == "" {
			referer = "-"
		}
		log.Debugf(`%s "%s %s" %s %d %s %s`,
			c.Request.RemoteAddr,
			c.Request.Method, c.Request.RequestURI, latency, c.Writer.Status(),
			referer, c.Request.Header)
	}
}
