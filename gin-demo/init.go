package main

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"strings"
	"time"
)

var rdb *redis.Client
var DB *gorm.DB

func init1() {
	log.SetFlags(log.Lshortfile)

	env := strings.ToLower(os.Getenv("envType"))
	addr := "r-2zeuu64hewjynuchp8pd.redis.rds.aliyuncs.com:6379"
	password := "BgS6q5PV7WECWMU3QZvuQCWqJjc@nU"
	switch env {
	case "test":
		addr = "r-2zeuu64hewjynuchp8.redis.rds.aliyuncs.com:6379"
		password = "BgS6q5PV7WECWMU3QZvuQCWqJjc@nU"
	case "product":
		addr = "r-2zetuvcr3uwmx3h12w.redis.rds.aliyuncs.com:6379"
		password = "Z32XQjPW6HKFfX0TFAsEeS+!wOBuY"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:               addr,
		Password:           password, // no password set
		DB:                 0,        // use default DB
		PoolSize:           500,
		IdleTimeout:        time.Second,
		IdleCheckFrequency: 10 * time.Second,
		MinIdleConns:       3,
		MaxRetries:         3, // 最大重试次数
		DialTimeout:        2 * time.Second,
	})

	initDB(env)
}

func initDB(env string) {
	dsn := "xw_mall_dev_rw:8Ufp%mUzg*bnkHMZVt8MwVdCdV&PzG@tcp(rm-2ze58oz13gl4089i37o.mysql.rds.aliyuncs.com:3306)/growth_center?charset=utf8&loc=Local&parseTime=True"
	switch env {
	case "test":
		dsn = "xw_mall_dev_rw:8Ufp%mUzg*bnkHMZVt8MwVdCdV&PzG@tcp(rm-2ze58oz13gl4089i37o.mysql.rds.aliyuncs.com:3306)/growth_center?charset=utf8&loc=Local&parseTime=True"
	case "product":
		dsn = "growth_center_rw:Mj2JkGCJM9U!Tntg0tnsX3I^JEETrf@tcp(rm-2ze6mlhf82p2a222r.mysql.rds.aliyuncs.com:3306)/growth_center?charset=utf8&loc=Local&parseTime=True"
	}
	maxConn := 50       // 最大连接数
	maxIdleConn := 25   // 最大空闲连接数
	connMaxLife := 3600 // 连接最长持续时间， 默认1小时，单位秒
	isLog := true       // 是否记录日志  日志级别为info
	logLevel := logger.Silent
	if isLog {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 表前缀
		Logger:         newLogger,
	})
	if err != nil {
		panic(err)
	}
	// 配置读写分离
	dbResolverCfg := dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn)},
		Replicas: []gorm.Dialector{},
		Policy:   dbresolver.RandomPolicy{},
	}
	err = DB.Use(
		dbresolver.Register(dbResolverCfg).
			SetConnMaxLifetime(time.Duration(connMaxLife) * time.Second).
			SetMaxIdleConns(maxIdleConn).
			SetMaxOpenConns(maxConn),
	)
	if err != nil {
		panic(err)
	}
}
