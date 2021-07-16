package models

import (
	"hellobeego/consts"
)

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
