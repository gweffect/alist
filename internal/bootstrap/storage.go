package bootstrap

import (
	"context"

	"github.com/gweffectx/safedav/internal/conf"
	"github.com/gweffectx/safedav/internal/db"
	"github.com/gweffectx/safedav/internal/model"
	"github.com/gweffectx/safedav/internal/op"
	"github.com/gweffectx/safedav/pkg/utils"
)

func LoadStorages() {
	storages, err := db.GetEnabledStorages()
	if err != nil {
		utils.Log.Fatalf("failed get enabled storages: %+v", err)
	}
	go func(storages []model.Storage) {
		for i := range storages {
			err := op.LoadStorage(context.Background(), storages[i])
			if err != nil {
				utils.Log.Errorf("failed get enabled storages: %+v", err)
			} else {
				utils.Log.Infof("success load storage: [%s], driver: [%s]",
					storages[i].MountPath, storages[i].Driver)
			}
		}
		conf.StoragesLoaded = true
	}(storages)
}
