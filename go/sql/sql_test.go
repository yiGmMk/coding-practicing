package dbcode

// import (
// 	"database/sql"
// 	"fmt"
// 	"testing"
// 	"time"

// 	. "github.com/smartystreets/goconvey/convey"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // database/sql
// func TestSql(t *testing.T) {
// 	Convey("sql", t, func() {
// 		db, err := sql.Open("", "")
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		defer func() {
// 			_ = db.Close()
// 		}()

// 		tx, err := db.Begin()
// 		if err != nil {
// 			t.Error("start transaction", err)
// 		}
// 		defer func() {
// 			_ = tx.Rollback()
// 		}()
// 	})
// }

// // gorm
// func TestGorm(t *testing.T) {
// 	db, err := gorm.Open(mysql.Open(""),
// 		&gorm.Config{
// 			SkipDefaultTransaction: true,
// 		})
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	sqlDB.SetConnMaxLifetime(time.Minute * 5)
// 	sqlDB.SetMaxIdleConns(3)
// 	sqlDB.SetMaxOpenConns(2)

// 	rows, err := sqlDB.Query("select * from user")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println(rows)
// }
