package sftp

import (
	"context"
	"os"
	"path"

	"github.com/gweffectx/safedav/internal/driver"
	"github.com/gweffectx/safedav/internal/errs"
	"github.com/gweffectx/safedav/internal/model"
	"github.com/gweffectx/safedav/pkg/utils"
	"github.com/pkg/sftp"
	log "github.com/sirupsen/logrus"
)

type SFTP struct {
	model.Storage
	Addition
	client *sftp.Client
}

func (d *SFTP) Config() driver.Config {
	return config
}

func (d *SFTP) GetAddition() driver.Additional {
	return &d.Addition
}

func (d *SFTP) Init(ctx context.Context) error {
	return d.initClient()
}

func (d *SFTP) Drop(ctx context.Context) error {
	if d.client != nil {
		_ = d.client.Close()
	}
	return nil
}

func (d *SFTP) List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
	log.Debugf("[sftp] list dir: %s", dir.GetPath())
	files, err := d.client.ReadDir(dir.GetPath())
	if err != nil {
		return nil, err
	}
	objs, err := utils.SliceConvert(files, func(src os.FileInfo) (model.Obj, error) {
		return d.fileToObj(src, dir.GetPath())
	})
	return objs, err
}

func (d *SFTP) Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	remoteFile, err := d.client.Open(file.GetPath())
	if err != nil {
		return nil, err
	}
	link := &model.Link{
		MFile: remoteFile,
	}
	return link, nil
}

func (d *SFTP) MakeDir(ctx context.Context, parentDir model.Obj, dirName string) error {
	return d.client.MkdirAll(path.Join(parentDir.GetPath(), dirName))
}

func (d *SFTP) Move(ctx context.Context, srcObj, dstDir model.Obj) error {
	return d.client.Rename(srcObj.GetPath(), path.Join(dstDir.GetPath(), srcObj.GetName()))
}

func (d *SFTP) Rename(ctx context.Context, srcObj model.Obj, newName string) error {
	return d.client.Rename(srcObj.GetPath(), path.Join(path.Dir(srcObj.GetPath()), newName))
}

func (d *SFTP) Copy(ctx context.Context, srcObj, dstDir model.Obj) error {
	return errs.NotSupport
}

func (d *SFTP) Remove(ctx context.Context, obj model.Obj) error {
	return d.remove(obj.GetPath())
}

func (d *SFTP) Put(ctx context.Context, dstDir model.Obj, stream model.FileStreamer, up driver.UpdateProgress) error {
	dstFile, err := d.client.Create(path.Join(dstDir.GetPath(), stream.GetName()))
	if err != nil {
		return err
	}
	defer func() {
		_ = dstFile.Close()
	}()
	err = utils.CopyWithCtx(ctx, dstFile, stream, stream.GetSize(), up)
	return err
}

var _ driver.Driver = (*SFTP)(nil)
