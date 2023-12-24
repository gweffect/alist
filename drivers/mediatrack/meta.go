package mediatrack

import (
	"github.com/gweffectx/safedav/internal/driver"
	"github.com/gweffectx/safedav/internal/op"
)

type Addition struct {
	AccessToken string `json:"access_token" required:"true"`
	ProjectID   string `json:"project_id"`
	driver.RootID
	OrderBy   string `json:"order_by" type:"select" options:"updated_at,title,size" default:"title"`
	OrderDesc bool   `json:"order_desc"`
}

var config = driver.Config{
	Name: "MediaTrack",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &MediaTrack{}
	})
}
