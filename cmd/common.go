package cmd

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/gweffectx/safedav/internal/bootstrap"
	"github.com/gweffectx/safedav/internal/bootstrap/data"
	"github.com/gweffectx/safedav/internal/db"
	"github.com/gweffectx/safedav/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func Init() {
	bootstrap.InitConfig()
	bootstrap.Log()
	bootstrap.InitDB()
	data.InitData()
	bootstrap.InitIndex()
}

func Release() {
	db.Close()
}

var pid = -1
var pidFile string

func initDaemon() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	_ = os.MkdirAll(filepath.Join(exPath, "daemon"), 0700)
	pidFile = filepath.Join(exPath, "daemon/pid")
	if utils.Exists(pidFile) {
		bytes, err := os.ReadFile(pidFile)
		if err != nil {
			log.Fatal("failed to read pid file", err)
		}
		id, err := strconv.Atoi(string(bytes))
		if err != nil {
			log.Fatal("failed to parse pid data", err)
		}
		pid = id
	}
}
