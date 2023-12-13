package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-house/consts"
	"github.com/go-house/library"
	dao "github.com/go-house/model/dao"
)

type UserAPI struct {
}

// Login 微信小程序登录
func (user *UserAPI) WXLogin(ctx *gin.Context) (any, *library.HError) {
	var (
		jscode string
	)

	res := make(map[string]any)

	jscode = ctx.Query("jscode")
	if jscode == "" {
		return nil, library.ErrorF(consts.ERR_PARAM_CODE, consts.ERR_PARAM_MEG)
	}
	resp, err := library.WxSDK.Code2Session(jscode)
	if err != nil || resp.ErrCode != 0 {
		return nil, library.ErrorF(consts.ERR_EXP_CODE, resp.ErrMSG)
	}
	userDao := dao.NewUserDao()
	userInfo := userDao.GetUserByOpenId(resp.Openid)
	isBind := false
	if userInfo == nil || userInfo.ID == 0 {
		// 注册用户
		userInfo.OpenID = resp.Openid
		userInfo.UnionID = resp.Unionid
		_, err := userDao.Insert(userInfo)
		if err != nil {
			return nil, err
		}
	}
	isBind = userInfo.Mobile != ""
	res["isBind"] = isBind
	token, err := library.GenerateJWT(userInfo.ID)
	res["token"] = token
	if err != nil {
		return nil, library.ErrorF(consts.ERR_LOGIN_CODE, consts.ERR_LOGIN_MEG)
	}
	// 判断是否绑定手机号等信息
	return res, nil
}
