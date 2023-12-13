package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-house/api/app"
	"github.com/go-house/library"
	"github.com/go-house/middleware"
)

// Rotuer 路由
func Router(r *gin.RouterGroup) {

	// 用户相关
	userAPI := new(app.UserAPI)
	userGroup := r.Group("/user")
	userGroup.GET("/wxlogin", request(userAPI.WXLogin))

	// 房间相关
	roomAPI := new(app.RoomAPI)
	roomGroup := r.Group("/room")
	roomGroup.GET("/list", middleware.CheckLogin(), request(roomAPI.List))

}

// request 封装返回
func request(f func(ctx *gin.Context) (any, *library.HError)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 错误拦截
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(500, library.RSystemErr())
				return
			}
		}()
		data, err := f(ctx)
		if err != nil {
			ctx.JSON(500, library.RERR(err))
			return
		}
		ctx.JSON(200, library.RSuccess(data))
	}
}
