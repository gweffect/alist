package common

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gweffectx/safedav/encrypt"
	"github.com/gweffectx/safedav/internal/model"
	"github.com/gweffectx/safedav/internal/net"
	"github.com/gweffectx/safedav/pkg/http_range"
	"github.com/gweffectx/safedav/pkg/utils"
)

func Proxy(w http.ResponseWriter, r *http.Request, link *model.Link, file model.Obj) error {
	if link.MFile != nil {
		defer link.MFile.Close()
		attachFileName(w, file)
		contentType := link.Header.Get("Content-Type")
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		http.ServeContent(w, r, file.GetName(), file.ModTime(), link.MFile)
		return nil
	} else if link.RangeReadCloser != nil {
		attachFileName(w, file)
		net.ServeHTTP(w, r, file.GetName(), file.ModTime(), file.GetSize(), link.RangeReadCloser.RangeRead)
		defer func() {
			_ = link.RangeReadCloser.Close()
		}()
		return nil
	} else if link.Concurrency != 0 || link.PartSize != 0 {
		attachFileName(w, file)
		size := file.GetSize()
		//var finalClosers model.Closers
		finalClosers := utils.EmptyClosers()
		header := net.ProcessHeader(r.Header, link.Header)
		rangeReader := func(ctx context.Context, httpRange http_range.Range) (io.ReadCloser, error) {
			down := net.NewDownloader(func(d *net.Downloader) {
				d.Concurrency = link.Concurrency
				d.PartSize = link.PartSize
			})
			req := &net.HttpRequestParams{
				URL:       link.URL,
				Range:     httpRange,
				Size:      size,
				HeaderRef: header,
			}
			rc, err := down.Download(ctx, req)
			finalClosers.Add(rc)
			return rc, err
		}
		net.ServeHTTP(w, r, file.GetName(), file.ModTime(), file.GetSize(), rangeReader)
		defer finalClosers.Close()
		return nil
	} else {
		//transparent proxy
		header := net.ProcessHeader(r.Header, link.Header)
		res, err := net.RequestHttp(context.Background(), r.Method, header, link.URL)
		if err != nil {
			return err
		}
		key := []byte("wumansgygoaescbc")
		safeReader := encrypt.NewDecryptReader(res.Body, key)
		defer safeReader.Close()
		contentRange := res.Header.Get("Content-Range")
		if contentRange != "" {
			contentRange = strings.Replace(contentRange, "bytes ", "", 1)
			rangeStartString := strings.Split(contentRange, "-")[0]
			if rangeStartString != "" {
				rangeStart, err := strconv.ParseInt(rangeStartString, 0, 64)
				if err == nil {
					if rangeStart > 0 {
						safeReader.SetOffset(rangeStart)
					}
				}
			}
		}
		fmt.Println("请求链接:" + link.URL)
		fmt.Println("响应Content-Range:" + res.Header.Get("Content-Range"))
		for h, v := range res.Header {
			w.Header()[h] = v
		}
		w.WriteHeader(res.StatusCode)
		if r.Method == http.MethodHead {
			return nil
		}

		_, err = io.Copy(w, safeReader)
		if err != nil {
			return err
		}
		return nil
	}
}
func attachFileName(w http.ResponseWriter, file model.Obj) {
	fileName := file.GetName()
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"; filename*=UTF-8''%s`, fileName, url.PathEscape(fileName)))
	w.Header().Set("Content-Type", utils.GetMimeType(fileName))
}
