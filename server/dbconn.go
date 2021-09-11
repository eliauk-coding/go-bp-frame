package server

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gobpframe/config"
	"gobpframe/utils/logger"
)

type DBConnect struct {
	Host     string
	Port     uint
	DBName   string
	SslMode  string
	Username string
	Password string
	TimeZone string
	db       *gorm.DB
}

const DefaultPort = 5432

var (
	dbConn      *DBConnect
	dbConnMutex sync.RWMutex
)

func DB() *gorm.DB {
	dbConnMutex.RLock()
	defer dbConnMutex.RUnlock()
	return dbConn.db
}

func DBConn() *DBConnect {
	dbConnMutex.RLock()
	defer dbConnMutex.RUnlock()
	return dbConn
}

func InitDB() {
	dbConnMutex.Lock()
	defer dbConnMutex.Unlock()

	if dbConn != nil {
		return
	}

	dsn, conn := loadDBConf()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Debug("database: ", dsn)
		logger.Fatal("open database failed:", err)
	}
	conn.db = db
	dbConn = conn

	maxIdle, maxOpen, connTTL := loadConnPoolConf()
	dbConn.SetConnectPool(maxIdle, maxOpen, connTTL)
}

func (c *DBConnect) Close() {
	dbConnMutex.Lock()
	defer dbConnMutex.Unlock()
	db, _ := c.db.DB()
	_ = db.Close()
}

func (c *DBConnect) SetConnectPool(maxIdle, maxOpen int, connTTL time.Duration) {
	db, _ := c.db.DB()
	db.SetMaxIdleConns(maxIdle)    // default 2. If n <= 0, no idle connections are retained.
	db.SetMaxOpenConns(maxOpen)    // default unlimited. If <= 0, there is no limit
	db.SetConnMaxLifetime(connTTL) // If <= 0, connections are not closed due to a connection's age.
}

func loadDBConf() (string, *DBConnect) {
	host := config.GetStr("Database.Host")
	port := config.GetUint("Database.Port")
	dbName := config.GetStr("Database.DBName")
	username := config.GetStr("Database.Username")
	password := config.GetStr("Database.Password")
	timeZone := config.GetStr("Database.TimeZone")
	sslMode := "disable"

	if host == "" || dbName == "" {
		logger.Fatal("database configuration is incorrect")
	}

	if port <= 0 {
		port = DefaultPort
	}

	if config.GetBool("Database.SslMode") {
		sslMode = "enable"
	}

	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=%s", host, port, dbName, sslMode)

	if username != "" {
		dsn += " user=" + username
	}

	if password != "" {
		dsn += " password=" + password
	}

	if timeZone != "" {
		dsn += " TimeZone=" + timeZone
	}

	dbConn = &DBConnect{
		Host:     host,
		Port:     port,
		DBName:   dbName,
		SslMode:  sslMode,
		Username: username,
		Password: password,
		TimeZone: timeZone,
	}
	return dsn, dbConn
}

func loadConnPoolConf() (maxIdle, maxOpen int, connTTL time.Duration) {
	maxIdle = config.GetInt("Database.ConnectPool.MaxIdleConns")
	maxOpen = config.GetInt("Database.ConnectPool.MaxOpenConns")
	connTTL = config.GetDuration("Database.ConnectPool.MaxConnLifetimeInMinutes") * time.Minute
	return maxIdle, maxOpen, connTTL
}
