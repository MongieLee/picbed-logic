package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var MysqlDB *gorm.DB

func InitMysqlDbConnector() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&parseTime=True&charset=utf8mb4",
		ProConfig.Mysql.Username,
		ProConfig.Mysql.Password,
		ProConfig.Mysql.Hostname,
		ProConfig.Mysql.Port,
		ProConfig.Mysql.Database,
	)
	fmt.Println(dsn)
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("连接失败", err)
	}
	return
}
