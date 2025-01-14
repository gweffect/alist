package pikpak_share

import (
	"github.com/gweffectx/safedav/internal/driver"
	"github.com/gweffectx/safedav/internal/op"
)

type Addition struct {
	driver.RootID
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
	ShareId  string `json:"share_id" required:"true"`
	SharePwd string `json:"share_pwd"`
}

var config = driver.Config{
	Name:        "PikPakShare",
	LocalSort:   true,
	NoUpload:    true,
	DefaultRoot: "",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &PikPakShare{}
	})
}
