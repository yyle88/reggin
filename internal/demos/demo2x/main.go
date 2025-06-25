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
	router.POST("/events/abc", warpginhandle.PX(abcHandle, makeResp[abcRet]))
	router.POST("/events/xyz", warpginhandle.PX(xyzHandle, makeResp[xyzRet]))

	done.Done(router.Run(":8080"))
}

type abcArg struct {
	Value string `json:"value"`
}

type abcRet struct {
	Num int `json:"num"`
}

func abcHandle(arg *abcArg) (*abcRet, error) {
	zaplog.SUG.Debugln("abc-arg:", neatjsons.S(arg))
	num, err := strconv.Atoi(arg.Value)
	if err != nil {
		return nil, erero.Wro(err)
	}
	ret := &abcRet{
		Num: num,
	}
	zaplog.SUG.Debugln("abc-ret:", neatjsons.S(ret))
	return ret, nil
}

type xyzArg struct {
	Num int `json:"num"`
}

type xyzRet struct {
	Value string `json:"value"`
}

func xyzHandle(arg *xyzArg) (*xyzRet, error) {
	zaplog.SUG.Debugln("xyz-arg:", neatjsons.S(arg))
	value := strconv.Itoa(arg.Num)
	ret := &xyzRet{
		Value: value,
	}
	zaplog.SUG.Debugln("xyz-ret:", neatjsons.S(ret))
	return ret, nil
}

type respType struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func makeResp[RES any](ctx *gin.Context, res *RES, err error) *respType {
	if err != nil {
		return &respType{
			Code: -1,
			Desc: err.Error(),
			Data: nil,
		}
	} else {
		return &respType{
			Code: 0,
			Desc: "SUCCESS",
			Data: res,
		}
	}
}
