/**
* @File: main.go
* @Author: wongxinjie
* @Date: 2019/10/6
 */
package main

import (
	"fmt"
	"time"

	"ip2region/cmd"
	"ip2region/db"
	"ip2region/model"
)

func main() {
	cmd.Execute()
}

func RunDatabase() {
	config := &db.Config{
		DatabaseURI: "root:@tcp(localhost:3306)/ip_region?charset=utf8mb4&parseTime=true",
	}

	database, err := db.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = database.AutoMigrate(&model.Passport{}).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	passports, err := database.ListPassport(10, 0, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", passports)
	}

	passport := &model.Passport{
		AppKey:    "131",
		AppSecret: "456",
		Status:    model.Using,
		ExpiredAt: time.Now(),
		CreatedAt: time.Now(),
	}
	err = database.CreatePassport(passport)
	if err != nil {
		fmt.Println(err)
	}

	passport, err = database.GetPassportByAppKey("123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", passport)
	}

	passport.AppSecret = "12345678"
	err = database.UpdatePassport(passport)
	if err != nil {
		fmt.Println(err)
	}

	passports, err = database.ListPassport(10, 0, -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", passports)
	}

}
