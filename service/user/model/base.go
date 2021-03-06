

package model

import (
	"github.com/Kydaa/bit/service/user/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", config.Db["db1"].Dsn)
	if err != nil {
		panic("数据库无法加载")
	}
}
