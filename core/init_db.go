package core

import (
	"blogx_server/global"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB{
	dc := global.Config.DB
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
	})
	if err != nil {
    logrus.Fatalf("数据库连接失败 %s", err)
  }
	fmt.Println(db, err)
  sqlDB, err := db.DB()
  sqlDB.SetMaxIdleConns(10)
  sqlDB.SetMaxOpenConns(100)
  sqlDB.SetConnMaxLifetime(time.Hour)
  logrus.Infof("数据库连接成功！")
  return db
}