package datasql

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/brunobotter/mercado-livre/configs/mapping"
	"github.com/brunobotter/mercado-livre/internal/domain/contract"
	"github.com/go-sql-driver/mysql"
)

var (
	instance     *Conn
	dbInstance   *sql.DB
	onceOB       sync.Once
	onceInstance sync.Once
	connErr      error
)

type Conn struct {
	db   *sql.DB
	user *userRepository
}

func Instance(cfg *mapping.Config) (contract.DataManager, error) {
	onceInstance.Do(func() {
		db, err := GetDB(cfg)
		if err != nil {
			connErr = errors.New(err.Error())
			return
		}

		instance = &Conn{db: db}
		instance.user = &userRepository{db, instance}

	})
	return instance, connErr
}

func GetDB(cfg *mapping.Config) (*sql.DB, error) {
	onceOB.Do(func() {
		mysqlCfg := getMySqlConfig(cfg)
		db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
		if err != nil {
			connErr = errors.New(err.Error())
			return
		}
		maxLifeTimeInMinutes, _ := time.ParseDuration(fmt.Sprintf("%vmin", cfg.DB.MaxLifeTimeInMinutes))
		db.SetConnMaxIdleTime(maxLifeTimeInMinutes)
		db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
		db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
		err = db.Ping()
		if err != nil {
			connErr = errors.New(err.Error())
			return
		}
		dbInstance = db
	})
	return dbInstance, connErr
}

func getMySqlConfig(cfg *mapping.Config) *mysql.Config {
	mysqlConfig := mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		DBName:               cfg.DB.Name,
		User:                 cfg.DB.User,
		Passwd:               cfg.DB.Pass,
		ParseTime:            true,
		AllowNativePasswords: true,
		Params:               cfg.DB.Params,
	}
	return &mysqlConfig
}

func (c *Conn) UserRepo() contract.UserRepository {
	return c.user
}
