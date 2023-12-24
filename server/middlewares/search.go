package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gweffectx/safedav/internal/conf"
	"github.com/gweffectx/safedav/internal/errs"
	"github.com/gweffectx/safedav/internal/setting"
	"github.com/gweffectx/safedav/server/common"
)

func SearchIndex(c *gin.Context) {
	mode := setting.GetStr(conf.SearchIndex)
	if mode == "none" {
		common.ErrorResp(c, errs.SearchNotAvailable, 500)
		c.Abort()
	} else {
		c.Next()
	}
}
