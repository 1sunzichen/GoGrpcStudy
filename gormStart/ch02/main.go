package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CompanyID 		int
	Company 	Company
}
type Company struct{
	ID int
	Name string
}
func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Info , // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)
	db, err:= gorm.Open(mysql.Open(dsn), &gorm.Config{Logger:newLogger, CreateBatchSize: 1000,})
	if err!=nil{
		panic(err)
	}

	// Migrate the schema 根据struct数据结构 创建出相同结构 表
	db.AutoMigrate(&User{})
	fmt.Println("mysql connect success")
	db.Create(&User{
		Name:"boddy2",
		Company: Company{
			ID:2,
		},
	})
	//var users = []User{{Name: "jinzhu_1"},{Name: "jinzhu_2"}, {Name: "jinzhu_10000"}}

	// 数量为 100
	//db.CreateInBatches(users, 100)
}