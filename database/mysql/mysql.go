package mysql

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mignfu_go_common/flags"
	"time"
)

var defaultGormConfig = &gorm.Config{
	PrepareStmt: true,
}

type ConnOptions struct {
	IP       string
	Port     int
	Username string
	Passwd   string
	DbName   string
}

func (o *ConnOptions) identity() (identity string) {
	identity = fmt.Sprintf("'%s'@'%s:%d'", o.Username, o.IP, o.Port)
	return
}

func (o *ConnOptions) dsn() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", o.Username, o.Passwd, o.IP, o.Port, o.DbName)
	return
}

func NewConnOptionsFromFlags() *ConnOptions {
	return &ConnOptions{
		IP:       flags.MysqlHost,
		Port:     flags.MysqlPort,
		Username: flags.MysqlUser,
		Passwd:   flags.MysqlPasswd,
		DbName:   flags.MysqlDB,
	}
}

func Connect(options *ConnOptions, config *gorm.Config, initCallback func(*gorm.DB) error) (db *gorm.DB, err error) {
	if config == nil {
		config = defaultGormConfig
	}

	db, err = gorm.Open(mysql.Open(options.dsn()), config)
	if err != nil {
		log.Errorf("Connect mysql %s failed, err: %v", options.identity(), err)
		return
	}
	log.Infof("Successfully connected to mysql %s", options.identity())

	if err = initCallback(db); err != nil {
		log.Errorf("Initialize mysql database %s failed, err: %v", options.identity(), err)
		return
	}
	log.Infof("Successfully initialize mysql")

	go monitorConnection(db, options, config)
	return
}

func monitorConnection(db *gorm.DB, options *ConnOptions, config *gorm.Config) {
	var err error
	var newDb *gorm.DB
	for {
		err = checkDBConnection(db)
		if err != nil {
			log.Errorf("Lost database connection, attempting to reconnect: %v", err)

			for {
				newDb, err = gorm.Open(mysql.Open(options.dsn()), config)
				if err != nil {
					log.Errorf("Reconnect failed, retrying in 2 seconds, err: %v", err)
					time.Sleep(2 * time.Second)
					continue
				}
				log.Info("Reconnected to MySQL successfully.")
				*db = *newDb
				break
			}
		}
		time.Sleep(2 * time.Second) // 每隔 2 秒检测一次连接
	}
}

func checkDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
