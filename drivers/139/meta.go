package _139

import (
	"github.com/gweffectx/safedav/internal/driver"
	"github.com/gweffectx/safedav/internal/op"
)

type Addition struct {
	//Account       string `json:"account" required:"true"`
	Authorization string `json:"authorization" type:"text" required:"true"`
	driver.RootID
	Type    string `json:"type" type:"select" options:"personal,family,personal_new" default:"personal"`
	CloudID string `json:"cloud_id"`
}

var config = driver.Config{
	Name:      "139Yun",
	LocalSort: true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Yun139{}
	})
}
