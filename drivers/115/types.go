package _115

import (
	"github.com/SheltonZhu/115driver/pkg/driver"
	"github.com/gweffectx/safedav/internal/model"
	"github.com/gweffectx/safedav/pkg/utils"
	"time"
)

var _ model.Obj = (*FileObj)(nil)

type FileObj struct {
	driver.File
}

func (f *FileObj) CreateTime() time.Time {
	return f.File.CreateTime
}

func (f *FileObj) GetHash() utils.HashInfo {
	return utils.NewHashInfo(utils.SHA1, f.Sha1)
}
