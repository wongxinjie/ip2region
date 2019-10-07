/**
* @File: passport.go
* @Author: wongxinjie
* @Date: 2019/10/7
 */
package api

import (
	"encoding/json"
	"net/http"
	"time"

	"ip2region/app"
)

type PassportResponse struct {
	AppKey    string    `json:"app_key"`
	AppSecret string    `json:"app_secret"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (a *API) CreatePassport(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	p, err := ctx.CreatePassport()
	if err != nil {
		return err
	}

	response, err := json.Marshal(&PassportResponse{
		AppKey:    p.AppKey,
		AppSecret: p.AppSecret,
		ExpiredAt: p.ExpiredAt,
	})
	if err != nil {
		return err
	}

	_, err = w.Write(response)
	return err
}
