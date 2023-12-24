package common

import (
	stdpath "path"

	"github.com/gweffectx/safedav/internal/conf"
	"github.com/gweffectx/safedav/internal/model"
	"github.com/gweffectx/safedav/internal/setting"
	"github.com/gweffectx/safedav/internal/sign"
)

func Sign(obj model.Obj, parent string, encrypt bool) string {
	if obj.IsDir() || (!encrypt && !setting.GetBool(conf.SignAll)) {
		return ""
	}
	return sign.Sign(stdpath.Join(parent, obj.GetName()))
}
