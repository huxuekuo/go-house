package library

import "github.com/go-house/consts"

type R struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func RSystemErr() *R {
	r := new(R)
	r.Code = consts.ERR_SYSTEM_CODE
	r.Msg = consts.ERR_SYSTEM_MEG
	return r
}

func RSuccess(data any) *R {
	r := new(R)
	r.Code = consts.SUCCESS_CODE
	r.Msg = consts.SUCCESS_MEG
	r.Data = data
	return r
}

func RERR(herr *HError) *R {
	r := new(R)
	r.Code = herr.Code
	r.Msg = herr.Message
	return r
}
