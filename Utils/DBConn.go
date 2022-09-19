package Utils

import (
	"MetaWebServer/DataReflect/Config"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	ConnQueue     []Connection
	CurrentSource string
)

type Connection struct {
	DB  *gorm.DB
	Key string
}

func CreateDBConns(conn Config.DBConn) *gorm.DB {
	var err error
	var DataBase *gorm.DB
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local`,
		conn.User, conn.Password, conn.Host, conn.Port, conn.DBName)
	DataBase, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("database connection error:%#v \n", err.Error())
		return nil
	} else {
		log.Printf("\n key :%#v\n host %#v \n connection established\n", conn.Key, conn.Host)
	}
	DataBase.DB().SetMaxIdleConns(conn.MaxConn)
	DataBase.DB().SetConnMaxLifetime(time.Duration(conn.MaxAlive))
	connection := &Connection{}
	connection.Key = conn.Key
	connection.DB = DataBase
	ConnQueue = append(ConnQueue, *connection)
	return DataBase
}

func GetConnByKey(key string) (*gorm.DB, error) {
	var err error
	var DataBase *gorm.DB
	if len(ConnQueue) == 0 {
		err = errors.New("no avalible connection,please keep a check on datasource.yaml")
		return nil, err
	}
	if key == "" {
		cacheMgr := CacheManager(&CacheService{})
		dbConfigCache, res := cacheMgr.Get(DB_G_KEY)
		if res {
			dbConfig := dbConfigCache.(Config.DBConfig)
			key = dbConfig.CurrentSource
		} else {
			err = errors.New("no avalible connection,please set default datasource")
			return nil, err
		}
	}
	for _, v := range ConnQueue {
		if v.Key == key {
			DataBase = v.DB
		}
	}
	return DataBase, err
}
