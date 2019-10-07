/**
* @File: passport.go
* @Author: wongxinjie
* @Date: 2019/10/7
 */
package db

import (
	"ip2region/model"

	"github.com/pkg/errors"
)

func (db *Database) ListPassport(limit, prevId int64, status int) ([]*model.Passport, error) {
	var passports []*model.Passport

	query := db.DB.Table("passport")
	if prevId != 0 && status != -1 {
		query = query.Where("id < ? and status = ?", prevId, status)
	} else if status != -1 {
		query = query.Where("status = ?", status)
	}

	err := query.Order("created_at desc").Limit(limit).Scan(&passports).Error

	return passports, errors.Wrap(err, "unable to list passport")
}

func (db *Database) CreatePassport(passport *model.Passport) error {
	return errors.Wrap(db.Create(passport).Error, "unable to create passport")
}

func (db *Database) GetPassportById(id int64) (*model.Passport, error) {
	var passport model.Passport

	err := db.First(&passport, id).Error
	return &passport, errors.Wrap(err, "unable to get passport by id")
}

func (db *Database) GetPassportByAppKey(appKey string) (*model.Passport, error) {
	var passport model.Passport

	err := db.First(&passport, model.Passport{AppKey: appKey}).Error
	return &passport, errors.Wrap(err, "unable to get passport by app_key")
}

func (db *Database) UpdatePassport(passport *model.Passport) error {
	err := db.Save(passport).Error
	return errors.Wrap(err, "unable to update passport")
}

func (db *Database) DeletePassport(id int64) error {
	err := db.Table(model.PassportTablename).Where("id = ?", id).Update(
		map[string]interface{}{"status": model.Deleted}).Error
	return errors.Wrap(err, "update passport error")
}
