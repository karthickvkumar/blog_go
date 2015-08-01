package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
func DBConn() gorm.DB {
	db, err := gorm.Open("mysql", "root:panda@/Blogger?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("Error:", err)
	}
	db.DB().SetMaxIdleConns(10)
	return db
		}



