package dao

import (
	"time"

	"github.com/go-house/consts"
	"github.com/go-house/library"
	"github.com/go-house/resource"
	"gorm.io/gorm"
)

// UserDao _
type UserDao struct {
	DB *gorm.DB
}

// User _
type User struct {
	ID         uint64 `gorm:"column:id"`
	UserName   string `gorm:"column:username"; json:"username"`
	Mobile     string `gorm:"column:mobile"`
	IsManage   uint8  `gorm:"column:is_manage"`
	OpenID     string `gorm:"column:openid"`
	UnionID    string `gorm:"column:unionid"`
	Deleted    uint8  `gorm:"column:deleted"`
	Createtime int64  `gorm:"column:createtime"`
	Updatetime int64  `gorm:"column:updatetime"`
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "app_user"
}

// NewUserDao _
func NewUserDao() *UserDao {
	instance := new(UserDao)
	instance.DB = resource.HouseDB
	return instance
}

// GetUserByOpenId 根据微信OpenID获取用户信息
func (user *UserDao) GetUserByOpenId(OpenID string) *User {
	userInfo := new(User)
	user.DB.Where("openid = ? AND deleted = 0", OpenID).First(&userInfo)
	return userInfo
}

// Insert 添加用户数据
func (user *UserDao) Insert(userInfo *User) (uint64, *library.HError) {
	nowUnix := time.Now().Unix()
	userInfo.Createtime = nowUnix
	userInfo.Updatetime = nowUnix
	result := user.DB.Create(userInfo)
	if result.Error != nil {
		return 0, library.ErrorF(consts.ERR_QUERY_RAL_CODE, consts.ERR_QUERY_RAL_MEG)
	}
	return userInfo.ID, nil
}
