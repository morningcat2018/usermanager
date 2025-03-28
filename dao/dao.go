package dao

import (
	"database/sql"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db    *gorm.DB
	SqlDB *sql.DB
)

func init() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	var err error
	Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if Db.Error != nil {
		panic("database error ")
	}
	SqlDB, err := Db.DB()
	if err != nil {
		panic("database error ")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	SqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	SqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	SqlDB.SetConnMaxLifetime(time.Hour)
}

// func GetDB() *gorm.DB {
// 	Db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	if Db.Error != nil {
// 		panic("database error ")
// 	}
// 	return Db
// }
