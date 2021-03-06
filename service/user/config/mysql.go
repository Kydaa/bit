/*
	加载配置
	author：Kyda
	date:06.18.2020
*/

package config

// DbConfig
type DbConfig struct {
	DriverName string
	Dsn string
	ShowSql bool
	ShowExecTime bool
	MaxIdle int
	MaxOpen int
}

var Db = map[string]DbConfig{
	"db1": {
		DriverName: "mysql",
		Dsn:        "root:123456@tcp(127.0.0.1:3306)/test2?charset=utf8&parseTime=True&loc=Local",
		ShowSql:    true,
		ShowExecTime:    false,
		MaxIdle:    10,
		MaxOpen:    200,
	},
}