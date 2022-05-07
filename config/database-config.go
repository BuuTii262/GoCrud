package config

import (
	"fmt"

	"exampl.com/goCrud/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func OpenDB() *gorm.DB {
	db_username := "root"
	db_host := "localhost"
	db_password := "123456"
	db_name := "gocurd"
	db_port := "3306"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_name)
	sql := mysql.Open(dsn)
	db, err := gorm.Open(sql, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("connect to Database")

	db.AutoMigrate(
		&model.User{},
		// &model.Money{},
	)

	DB = db
	return db

}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		println(err.Error())
	}
	sqlDB.Close()
}
