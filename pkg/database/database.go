package database

import (
	"docman/cfg"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Inst *gorm.DB

func Connect() error {
	var dbConf = cfg.Config.Database
	fmt.Printf("%v", dbConf)
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	Inst = db
	return err
}
