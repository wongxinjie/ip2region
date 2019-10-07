/**
* @File: model.go
* @Author: wongxinjie
* @Date: 2019/10/6
*/
package model

import (
	"crypto/rand"
	"time"
)

type Id []byte

func NewRequestId() Id {
	r := make(Id, 20)
	if _, err := rand.Read(r); err != nil {
		panic(err)
	}
	return r
}

type RequestId struct {
	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}