package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"goapp/internal/pkg/config"
	"goapp/internal/pkg/models/sys"
	"goapp/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

func InitDB(config *config.Config) {
	var gdb *gorm.DB
	var err error
	if config.Gorm.DBType=="mysql"{
		config.Gorm.DSN=config.MySQL.DSN()
	} else if config.Gorm.DBType=="sqlite3"{
		config.Gorm.DSN=config.Sqlite3.DSN()
	}
	gdb, err = gorm.Open(config.Gorm.DBType, config.Gorm.DSN)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true) 
	if config.Gorm.Debug {
		gdb.LogMode(true)                           
		gdb.SetLogger(log.New(os.Stdout, "\r\n", 0)) 
	}
	gdb.DB().SetMaxIdleConns(config.Gorm.MaxIdleConns)                                
	gdb.DB().SetMaxOpenConns(config.Gorm.MaxOpenConns)                                
	gdb.DB().SetConnMaxLifetime(time.Duration(config.Gorm.MaxLifetime) * time.Second) 
  db.DB=gdb
}

func Migration() {
	fmt.Println(db.DB.AutoMigrate(new(sys.Menu)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.Admins)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.RoleMenu)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.Role)).Error)
	fmt.Println(db.DB.AutoMigrate(new(sys.AdminsRole)).Error)
}
