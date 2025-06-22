package dao

import (
	"discuss/internal/config"
	"discuss/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var GormDb *gorm.DB

func init() {
	conf := config.GetConfig().MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DatabaseName,
	)

	var err error
	GormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}

	err = GormDb.AutoMigrate(&model.UserInfo{})
	if err != nil {
		log.Fatalf("Failed to auto migrate UserInfo model: %v", err.Error())
	}
}
