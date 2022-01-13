package common

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	logger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

type DB struct {
	DbMySql *gorm.DB
}

var (
	onceDbMysql sync.Once
	instanceDB  *DB
)

func GetInstanceDb() *gorm.DB {
	onceDbMysql.Do(func() {
		mysqlInfo := FileConfig.Database.MySql
		logs := fmt.Sprintf("[INFO] Connected to DB TYPE = %s ", mysqlInfo.Host)
		//dataSourceName := "root:@tcp(localhost:3306)/apps_db?parseTime=True&loc=Asia%2FJakarta"
		dbConfig := mysqlInfo.Username + ":" + mysqlInfo.Password + "@tcp(" + mysqlInfo.Host + ":" + mysqlInfo.Port + ")/" + mysqlInfo.Name + "?parseTime=True&loc=Asia%2FJakarta"
		sqlDB, err := sql.Open("mysql", dbConfig)
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to DB with err %s. Config=%s", err.Error(), mysqlInfo.Host)
			log.Fatalln(logs)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(20)
		sqlDB.SetConnMaxLifetime(10 * time.Minute)
		dialect := mysql.New(mysql.Config{Conn: sqlDB})
		loggerLevel := logger.Error

		dbConnection, err := gorm.Open(dialect, &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      loggerLevel,
				},
			),
		})
		if err != nil {
			logs = fmt.Sprintf("[ERROR] Failed to connect to DB with err %s. Config=%s", err.Error(), mysqlInfo.Host)
			log.Fatalln(logs)
		}
		fmt.Println(logs)
		instanceDB = &DB{DbMySql: dbConnection}
	})
	return instanceDB.DbMySql
}
