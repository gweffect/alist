package smb

import (
	"github.com/gweffectx/safedav/internal/driver"
	"github.com/gweffectx/safedav/internal/op"
)

type Addition struct {
	driver.RootPath
	Address   string `json:"address" required:"true"`
	Username  string `json:"username" required:"true"`
	Password  string `json:"password"`
	ShareName string `json:"share_name" required:"true"`
}

var config = driver.Config{
	Name:        "SMB",
	LocalSort:   true,
	OnlyLocal:   true,
	DefaultRoot: ".",
	NoCache:     true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &SMB{}
	})
}
