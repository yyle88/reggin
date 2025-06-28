package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/reggin/warpginhandle"
	"github.com/yyle88/zaplog"
)

func main() {
	router := gin.Default()
	router.POST("/api/aaa", warpginhandle.RX(aaaHandle, wrongResp))
	router.POST("/api/bbb", warpginhandle.FX(bbbHandle, wrongResp))

	done.Done(router.Run(":8080"))
}

type aaaArg struct {
	Value string `json:"value"`
}

type aaaRes struct {
	Num int `json:"num"`
}

func aaaHandle(arg *aaaArg) (*aaaRes, error) {
	zaplog.SUG.Debugln("aaa-arg:", neatjsons.S(arg))
	num, err := strconv.Atoi(arg.Value)
	if err != nil {
		return nil, erero.Wro(err)
	}
	ret := &aaaRes{
		Num: num,
	}
	zaplog.SUG.Debugln("aaa-res:", neatjsons.S(ret))
	return ret, nil
}

type bbbArg struct {
	Num int `json:"num"`
}

type bbbRet struct {
	Value string `json:"value"`
}

func bbbHandle(ctx *gin.Context, arg *bbbArg) (*bbbRet, error) {
	zaplog.SUG.Debugln("bbb-arg:", neatjsons.S(arg))
	value := strconv.Itoa(arg.Num)
	ret := &bbbRet{
		Value: value,
	}
	zaplog.SUG.Debugln("bbb-res:", neatjsons.S(ret))
	return ret, nil
}

type causeType struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

func wrongResp(ctx *gin.Context, cause error) *causeType {
	return &causeType{
		Code: -1,
		Desc: cause.Error(),
	}
}
