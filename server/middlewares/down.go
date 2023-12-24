package middlewares

import (
	"strings"

	"github.com/gweffectx/safedav/internal/conf"
	"github.com/gweffectx/safedav/internal/setting"

	"github.com/gin-gonic/gin"
	"github.com/gweffectx/safedav/internal/errs"
	"github.com/gweffectx/safedav/internal/model"
	"github.com/gweffectx/safedav/internal/op"
	"github.com/gweffectx/safedav/internal/sign"
	"github.com/gweffectx/safedav/pkg/utils"
	"github.com/gweffectx/safedav/server/common"
	"github.com/pkg/errors"
)

func Down(c *gin.Context) {
	rawPath := parsePath(c.Param("path"))
	c.Set("path", rawPath)
	meta, err := op.GetNearestMeta(rawPath)
	if err != nil {
		if !errors.Is(errors.Cause(err), errs.MetaNotFound) {
			common.ErrorResp(c, err, 500, true)
			return
		}
	}
	c.Set("meta", meta)
	// verify sign
	if needSign(meta, rawPath) {
		s := c.Query("sign")
		err = sign.Verify(rawPath, strings.TrimSuffix(s, "/"))
		if err != nil {
			common.ErrorResp(c, err, 401)
			c.Abort()
			return
		}
	}
	c.Next()
}

// TODO: implement
// path maybe contains # ? etc.
func parsePath(path string) string {
	return utils.FixAndCleanPath(path)
}

func needSign(meta *model.Meta, path string) bool {
	if setting.GetBool(conf.SignAll) {
		return true
	}
	if common.IsStorageSignEnabled(path) {
		return true
	}
	if meta == nil || meta.Password == "" {
		return false
	}
	if !meta.PSub && path != meta.Path {
		return false
	}
	return true
}
