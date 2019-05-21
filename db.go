package public_dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)


var DB *gorm.DB
func init() {
	url :="root:123456@/mydb01?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open("mysql", url)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	DB.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	DB.DB().SetConnMaxLifetime(time.Hour)
}