/**
* @File: client.go
* @Author: wongxinjie
* @Date: 2019/10/13
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func simpleRequest() {
	resp, err := http.Get("http://httpbin.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func postWithContent() error {
	type R struct {
		TeacherID uint64 `json:"teacher_id"`
		Name      string `json:"name"`
		Mobile    string `json:"mobile"`
	}

	r := R{
		TeacherID: 1234567,
		Name:      "wongxinjie",
		Mobile:    "18825111143",
	}
	payload, err := json.Marshal(r)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://httpbin.org/anything", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func requestWithHeaders() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://httpbin.org/headers", nil)
	if err != nil {
		return err
	}
	req.Header.Add("X-Token", "123455")
	req.Header.Add("X-Request-By", "wongxinjie")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil

	//resp, err := client.Get("http://httpbin.org/ip")
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}
	//
	//type IPResponse struct {
	//	Origin string `json:"origin"`
	//}
	//ip := new(IPResponse)
	//
	//err = json.Unmarshal(body, &ip)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	//fmt.Println(ip)
	//return nil
}

func main() {
	postWithContent()
}
