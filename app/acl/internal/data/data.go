package data

import (
	"time"
)

import (
	"github.com/dubbogo/gost/log/logger"

	"github.com/google/wire"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

import (
	"github.com/yrcs/nicehouse/app/acl/internal/data/po"
)

var ProviderSet = wire.NewSet(NewDB, NewData, NewRoleRepo)

type Data struct {
	db *gorm.DB
}

func NewDB(conf map[string]any) *gorm.DB {
	logger.Info("module: nicehouse-acl/data/NewDB")

	conf = conf["database"].(map[string]any)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf["source"].(string), // DSN (data source name)
		DefaultStringSize:         256,                     // Default length of a string type field
		DisableDatetimePrecision:  true,                    // Disable datetime precision, not supported before MySQL 5.6
		SkipInitializeWithVersion: false,                   // Automatic configuration based on current MySQL version
	}), &gorm.Config{
		Logger:                                   gl.Default.LogMode(gl.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true, NoLowerCase: true},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
	})
	if err != nil {
		logger.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色'").AutoMigrate(&po.Role{}); err != nil {
		logger.Fatal(err)
	}

	pool := conf["connection-pool"].(map[string]any)
	maxIdleConns := pool["max_idle_conns"].(int)
	maxOpenConns := pool["max_open_conns"].(int)
	maxIdleTime := pool["max_idle_time"].(int)
	maxLifeTime := pool["max_life_time"].(int)

	sqlDB, _ := db.DB()
	// Set the maximum number of idle connections in the connection pool
	sqlDB.SetMaxIdleConns(maxIdleConns)
	// Set the maximum number of open database connections
	sqlDB.SetMaxOpenConns(maxOpenConns)
	// Set the maximum amount of time a connection can be left idle
	sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(maxIdleTime))
	// Set the maximum duration of time a connection can be reused
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(maxLifeTime))
	return db
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	d := &Data{db}
	cleanup := func() {
		logger.Info("closing the data resources")

		sqlDB, _ := db.DB()
		if err := sqlDB.Close(); err != nil {
			logger.Error(err)
		}
	}
	return d, cleanup, nil
}
