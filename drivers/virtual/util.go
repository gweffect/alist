package virtual

import (
	"time"

	"github.com/gweffectx/safedav/internal/model"
	"github.com/gweffectx/safedav/pkg/utils/random"
)

func (d *Virtual) genObj(dir bool) model.Obj {
	obj := &model.Object{
		Name:     random.String(10),
		Size:     0,
		IsFolder: true,
		Modified: time.Now(),
	}
	if !dir {
		obj.Size = random.RangeInt64(d.MinFileSize, d.MaxFileSize)
		obj.IsFolder = false
	}
	return obj
}
