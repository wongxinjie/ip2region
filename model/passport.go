/**
* @File: password.go
* @Author: wongxinjie
* @Date: 2019/10/7
 */
package model

import (
	"fmt"
	"math/rand"
	"time"

	"ip2region/common"
)

const (
	Using = iota
	Expired
	Deleted
)

const PassportTablename = "passport"

type Passport struct {
	Id        int64  `gorm:"primary_key;auto_increment"`
	AppKey    string `gorm:"unique;not null;size:64"`
	AppSecret string `gorm:"size:256"`
	Status    int8
	ExpiredAt time.Time
	CreatedAt time.Time
}

func (p Passport) TableName() string {
	return "passport"
}

func (p *Passport) String() string {
	return fmt.Sprintf("%s{Id: %d, AppKey: %s}", p.TableName(), p.Id, p.AppKey)
}

func (p *Passport) GenerateAppKey() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d%d", rand.Int(), time.Now().Unix())
}

func (p *Passport) GenerateAppSecret() string {
	return common.RandomString(64)
}
