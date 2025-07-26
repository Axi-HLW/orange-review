package mysql

import (
	"github.com/yzc/orange-review/app/review/biz/model"
	"github.com/yzc/orange-review/app/review/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(
		&model.ReviewInfo{},
		&model.ReviewAppealInfo{},
		&model.ReviewReplyInfo{},
	)
}
