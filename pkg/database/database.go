package database

import (
	"docman/cfg"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Inst *gorm.DB

func Connect() error {
	var dbConf = cfg.Config.Database
	// pgsql
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConf.Host, dbConf.User, dbConf.Password, dbConf.Name, dbConf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if cfg.Config.Server.Env == "dev" {
		db = db.Debug()
	}
	Inst = db
	return err
}
