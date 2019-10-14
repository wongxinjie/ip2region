/**
* @File: ip2region.go
* @Author: wongxinjie
* @Date: 2019/10/6
 */
package api

import (
	"net/http"

	"ip2region/app"
)

func (a *API) GetIPRegion(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte(`{"message": "ok"}`))
	return err
}
