package db

import (
	"fmt"

	"github.com/kafka-embracetheday/cuddly-octo-waffle/common/db/mysql"
	"github.com/kafka-embracetheday/cuddly-octo-waffle/common/logger"
	"github.com/kafka-embracetheday/cuddly-octo-waffle/config"
)

func Init() {
	cfg := config.Get()

	gormCfg := &mysql.Gorm{
		Dsn:             cfg.Mysql.Dsn,
		MaxOpenConns:    cfg.Mysql.MaxOpenConns,
		MaxIdleConns:    cfg.Mysql.MaxIdleConns,
		ConnMaxLifetime: cfg.Mysql.ConnMaxLifetime,
	}
	if err := gormCfg.Init(); err != nil {
		logger.Fatal(fmt.Sprintf("init mysql failed, dsn: %v, err: %v", cfg.Mysql.Dsn, err))
		panic(err)
	}
}
