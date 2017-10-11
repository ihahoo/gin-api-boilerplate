package db

import (
	"database/sql"
	"github.com/ihahoo/go-api-lib/config"
	"github.com/ihahoo/go-api-lib/log"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	user := config.GetString("db.options.user")
	password := config.GetString("db.options.password")
	host := config.GetString("db.options.host")
	port := config.GetString("db.options.port")
	dbname := config.GetString("db.options.dbname")
	connectTimeout := config.GetString("db.options.connectTimeout")

	dbIns, err := ConnectDB("dbname=" + dbname + " user=" + user + " password=" + password + " host=" + host + " port=" + port + " connect_timeout=" + connectTimeout + " sslmode=disable")
	if err != nil {
		log.GetLog().Fatal(err)
	}
	db = dbIns
}

// ConnectDB 连接数据库
func ConnectDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// GetDB 获得db客户端实例
func Client() *sql.DB {
	return db
}
