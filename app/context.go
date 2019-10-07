/**
* @File: context.go
* @Author: wongxinjie
* @Date: 2019/10/6
 */
package app

import (
	"github.com/sirupsen/logrus"

	"ip2region/db"
)

type Context struct {
	Logger        logrus.FieldLogger
	RemoteAddress string
	Database      *db.Database
}

func (ctx *Context) WithLogger(logger logrus.FieldLogger) *Context {
	c := *ctx
	c.Logger = logger
	return &c
}

func (ctx *Context) WithRemoteAddress(addr string) *Context {
	c := *ctx
	c.RemoteAddress = addr
	return &c
}
