package bootstrap

import (
	"github.com/gweffectx/safedav/internal/conf"
	"github.com/gweffectx/safedav/internal/fs"
	"github.com/gweffectx/safedav/internal/offline_download/tool"
	"github.com/xhofe/tache"
)

func InitTaskManager() {
	fs.UploadTaskManager = tache.NewManager[*fs.UploadTask](tache.WithWorks(conf.Conf.Tasks.Upload.Workers), tache.WithMaxRetry(conf.Conf.Tasks.Upload.MaxRetry))
	fs.CopyTaskManager = tache.NewManager[*fs.CopyTask](tache.WithWorks(conf.Conf.Tasks.Copy.Workers), tache.WithMaxRetry(conf.Conf.Tasks.Copy.MaxRetry))
	tool.DownloadTaskManager = tache.NewManager[*tool.DownloadTask](tache.WithWorks(conf.Conf.Tasks.Download.Workers), tache.WithMaxRetry(conf.Conf.Tasks.Download.MaxRetry))
	tool.TransferTaskManager = tache.NewManager[*tool.TransferTask](tache.WithWorks(conf.Conf.Tasks.Transfer.Workers), tache.WithMaxRetry(conf.Conf.Tasks.Transfer.MaxRetry))
}
