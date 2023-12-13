package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-house/library"
	"github.com/go-house/model/dao"
)

type RoomAPI struct{}

func (r *RoomAPI) List(ctx *gin.Context) (any, *library.HError) {
	roomDao := dao.NewRoomDao()
	roomList, err := roomDao.List(1, 1, 1, 1)
	if err != nil {
		return nil, err
	}
	return roomList, nil
}
