package mysql

import (
	"github.com/yzc/orange-review/app/review_server/biz/model"
	"github.com/yzc/orange-review/app/review_server/conf"

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

	err = Migrate()
	if err != nil {
		panic(err)
	}
}

// 自动迁移表结构
func Migrate() (err error) {
	err = DB.AutoMigrate(
		&model.ReviewInfo{},
		&model.ReviewReplyInfo{},
		&model.ReviewAppealInfo{},
	)
	return err
}
