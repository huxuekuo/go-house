package library

import (
	"github.com/go-house/resource"
	"github.com/medivhzhan/weapp/v3"
	"github.com/medivhzhan/weapp/v3/auth"
)

type LocalWxSDK struct {
	Appid     string
	Appsecret string
}

var WxSDK *LocalWxSDK

func InitWxSDK(conf map[string]string) {
	resource.WxSDK = weapp.NewClient(conf["wx_app_id"], conf["wx_app_secret"])
	WxSDK = new(LocalWxSDK)
	WxSDK.Appid = conf["wx_app_id"]
	WxSDK.Appsecret = conf["wx_app_secret"]
}

func (local *LocalWxSDK) Code2Session(jscode string) (*auth.Code2SessionResponse, error) {
	// 获取用户OpenID
	code2s := new(auth.Code2SessionRequest)
	code2s.Appid = local.Appid
	code2s.Secret = local.Appsecret
	code2s.JsCode = jscode
	code2s.GrantType = "grant_type"
	resp, err := resource.WxSDK.NewAuth().Code2Session(code2s)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
