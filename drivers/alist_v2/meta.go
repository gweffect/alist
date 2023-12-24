package alist_v2

import (
	"github.com/gweffectx/safedav/internal/driver"
	"github.com/gweffectx/safedav/internal/op"
)

type Addition struct {
	driver.RootPath
	Address     string `json:"url" required:"true"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

var config = driver.Config{
	Name:        "AList V2",
	LocalSort:   true,
	NoUpload:    true,
	DefaultRoot: "/",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &AListV2{}
	})
}
