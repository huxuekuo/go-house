package library

import (
	"fmt"

	"github.com/go-house/resource"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(conf map[string]string) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", conf["username"], conf["password"], conf["host"], conf["port"], conf["dbname"], conf["ext"])
	resource.HouseDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		panic(err)
	}
}
