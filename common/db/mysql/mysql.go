package mysql

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Gorm struct {
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

var GormDB *DB

type DB struct {
	DB *gorm.DB
}

func (g *Gorm) Init() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
			LogLevel:                  logger.Warn,
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: g.Dsn,
	}), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(g.MaxOpenConns)
	sqlDB.SetMaxIdleConns(g.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(g.ConnMaxLifetime) * time.Minute)
	GormDB = &DB{
		DB: db,
	}
	return nil
}
