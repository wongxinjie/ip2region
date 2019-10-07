/**
* @File: passport.go
* @Author: wongxinjie
* @Date: 2019/10/7
 */
package app

import (
	"time"

	"ip2region/model"
)

func (ctx *Context) CreatePassport() (*model.Passport, error) {
	passport := &model.Passport{
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * time.Duration(365*24)),
		Status:    model.Using,
	}
	passport.AppKey = passport.GenerateAppKey()
	passport.AppSecret = passport.GenerateAppSecret()

	return passport, ctx.Database.CreatePassport(passport)
}

func (ctx *Context) ListPassport(limit, prevId int64, status int) ([]*model.Passport, error) {
	return ctx.Database.ListPassport(limit, prevId, status)
}

func (ctx *Context) DeletePassport(id int64) error {
	return ctx.Database.DeletePassport(id)
}
