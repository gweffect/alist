package alist_v3

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gweffectx/safedav/drivers/base"
	"github.com/gweffectx/safedav/internal/op"
	"github.com/gweffectx/safedav/pkg/utils"
	"github.com/gweffectx/safedav/server/common"
	log "github.com/sirupsen/logrus"
)

func (d *AListV3) login() error {
	var resp common.Resp[LoginResp]
	_, err := d.request("/auth/login", http.MethodPost, func(req *resty.Request) {
		req.SetResult(&resp).SetBody(base.Json{
			"username": d.Username,
			"password": d.Password,
		})
	})
	if err != nil {
		return err
	}
	d.Token = resp.Data.Token
	op.MustSaveDriverStorage(d)
	return nil
}

func (d *AListV3) request(api, method string, callback base.ReqCallback, retry ...bool) ([]byte, error) {
	url := d.Address + "/api" + api
	req := base.RestyClient.R()
	req.SetHeader("Authorization", d.Token)
	if callback != nil {
		callback(req)
	}
	res, err := req.Execute(method, url)
	if err != nil {
		return nil, err
	}
	log.Debugf("[alist_v3] response body: %s", res.String())
	if res.StatusCode() >= 400 {
		return nil, fmt.Errorf("request failed, status: %s", res.Status())
	}
	code := utils.Json.Get(res.Body(), "code").ToInt()
	if code != 200 {
		if (code == 401 || code == 403) && !utils.IsBool(retry...) {
			err = d.login()
			if err != nil {
				return nil, err
			}
			return d.request(api, method, callback, true)
		}
		return nil, fmt.Errorf("request failed,code: %d, message: %s", code, utils.Json.Get(res.Body(), "message").ToString())
	}
	return res.Body(), nil
}

func (d *AListV3) requestWithTimeout(api, method string, callback base.ReqCallback, timeout time.Duration, retry ...bool) ([]byte, error) {
	url := d.Address + "/api" + api
	client := base.NewRestyClient().SetTimeout(timeout)
	req := client.R()
	req.SetHeader("Authorization", d.Token)
	if callback != nil {
		callback(req)
	}
	res, err := req.Execute(method, url)
	if err != nil {
		return nil, err
	}
	log.Debugf("[alist_v3] response body: %s", res.String())
	if res.StatusCode() >= 400 {
		return nil, fmt.Errorf("request failed, status: %s", res.Status())
	}
	code := utils.Json.Get(res.Body(), "code").ToInt()
	if code != 200 {
		if (code == 401 || code == 403) && !utils.IsBool(retry...) {
			err = d.login()
			if err != nil {
				return nil, err
			}
			return d.requestWithTimeout(api, method, callback, timeout, true)
		}
		return nil, fmt.Errorf("request failed,code: %d, message: %s", code, utils.Json.Get(res.Body(), "message").ToString())
	}
	return res.Body(), nil
}
