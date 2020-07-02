package models

import (
	"lxtkj/hellobeego/consts"
)

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code consts.JsonResultCode `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
}

type ListJsonResult struct {
	Code  consts.JsonResultCode `json:"code"`
	Msg   string                `json:"msg"`
	Count int64                 `json:"count"`
	Data  interface{}           `json:"data"`
}