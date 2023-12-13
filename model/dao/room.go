package dao

import (
	"github.com/go-house/consts"
	"github.com/go-house/library"
	"github.com/go-house/resource"
	"gorm.io/gorm"
)

type RoomDao struct {
	DB *gorm.DB
}

type Room struct {
	ID         uint64 `gorm:"column:id"`
	RoomNum    string `gorm:"column:room_num"`
	Address    string `gorm:"column:address"`
	Pic        string `gorm:"column:pic"`
	Pics       string `gorm:"column:pics"`
	State      uint8  `gorm:"column:state"`
	ReserveNum int64  `gorm:"column:reserve_num"`
	Password   string `gorm:"column:password"`
	DescStr    string `gorm:"column:desc_str"`
	Deleted    uint8  `gorm:"column:deleted"`
	Createtime int64  `gorm:"column:createtime"`
	Updatetime int64  `gorm:"column:updatetime"`
}

func NewRoomDao() *RoomDao {
	instance := new(RoomDao)
	instance.DB = resource.HouseDB
	return instance
}

func (r *RoomDao) List(startUnix, endUnix, page, size int64) (*[]Room, *library.HError) {
	var RoomList []Room
	err := r.DB.Where("startUinx = ? AND endUnix = ? LIMIT ?,?", startUnix, endUnix, page, size).Find(&RoomList)
	if err.Error != nil {
		return nil, library.ErrorF(consts.ERR_QUERY_RAL_CODE, consts.ERR_QUERY_RAL_MEG)
	}
	return &RoomList, nil
}
