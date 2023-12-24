package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gweffectx/safedav/internal/conf"
)

func ForceHttps(c *gin.Context) {
	if c.Request.TLS == nil {
		host := c.Request.Host
		// change port to https port
		host = strings.Replace(host, fmt.Sprintf(":%d", conf.Conf.Scheme.HttpPort), fmt.Sprintf(":%d", conf.Conf.Scheme.HttpsPort), 1)
		c.Redirect(302, "https://"+host+c.Request.RequestURI)
		c.Abort()
		return
	}
	c.Next()
}
