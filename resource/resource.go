package resource

import (
	"github.com/medivhzhan/weapp/v3"
	"gorm.io/gorm"
)

// HouseDB 酒店Mysql
var HouseDB *gorm.DB

// WxSDK 微信小程序SDK
var WxSDK *weapp.Client
