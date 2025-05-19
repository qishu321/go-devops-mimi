package common

import (
	"fmt"

	"go-devops-mimi/server/config"
	"go-devops-mimi/server/model/cmdb"
	"go-devops-mimi/server/model/example"
	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/model/nav"
	"go-devops-mimi/server/model/system"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局数据库对象
var DB *gorm.DB

// 初始化数据库
func InitDB() {
	switch config.Conf.Database.Driver {
	case "mysql":
		DB = ConnMysql()
	case "sqlite3":
		DB = ConnSqlite()
	}
	dbAutoMigrate()
}

// 自动迁移表结构
func dbAutoMigrate() {
	_ = DB.AutoMigrate(
		&system.User{},
		&system.Role{},
		&system.Group{},
		&system.Menu{},
		&system.Api{},
		&system.OperationLog{},
		&example.CloudAccount{},
		&cmdb.Nodes{},
		&cmdb.NodeGroup{},
		&exec.ScriptLibrary{},
		&exec.ScriptLog{},
		&exec.Script{},
		&exec.Transfer{},
		&exec.Task{},
		&exec.TaskManage{},
		&exec.TaskLog{},
		&exec.TaskManageLog{},
		&exec.ManageLog{},
		&exec.Cron{},
		&exec.CronLog{},
		&nav.Nav{},
		&nav.Link{},
	)
}

func ConnSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.Conf.Database.Source), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		Log.Panicf("failed to connect sqlite3: %v", err)
	}
	dbObj, err := db.DB()
	if err != nil {
		Log.Panicf("failed to get sqlite3 obj: %v", err)
	}
	// 参见： https://github.com/glebarez/sqlite/issues/52
	dbObj.SetMaxOpenConns(1)
	return db
}

func ConnMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	// 隐藏密码
	showDsn := fmt.Sprintf(
		"%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		// 打印sql日志
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		Log.Panicf("初始化mysql数据库异常: %v", err)
	}
	// 开启mysql日志
	if config.Conf.Mysql.LogMode {
		db.Debug()
	}
	Log.Infof("初始化mysql数据库完成! dsn: %s", showDsn)
	return db
}
