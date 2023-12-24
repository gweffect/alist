package alias

import (
	"github.com/gweffectx/safedav/internal/driver"
	"github.com/gweffectx/safedav/internal/op"
)

type Addition struct {
	// Usually one of two
	// driver.RootPath
	// define other
	Paths string `json:"paths" required:"true" type:"text"`
}

var config = driver.Config{
	Name:        "Alias",
	LocalSort:   true,
	NoCache:     true,
	NoUpload:    true,
	DefaultRoot: "/",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Alias{}
	})
}
