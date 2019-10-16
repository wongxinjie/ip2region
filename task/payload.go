/**
* @File: payload.go
* @Author: wongxinjie
* @Date: 2019/10/17
 */
package task

type Payload struct {
	Id    int64  `json:"id"`
	Token string `json:"token"`
	Data  string `json:"data"`
}
