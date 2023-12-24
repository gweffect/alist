package data

import "github.com/gweffectx/safedav/cmd/flags"

func InitData() {
	initUser()
	initSettings()
	if flags.Dev {
		initDevData()
		initDevDo()
	}
}
