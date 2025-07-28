package dal

import (
	"github.com/yzc/orange-review/app/review/biz/dal/es"
	"github.com/yzc/orange-review/app/review/biz/dal/mysql"
	// "github.com/yzc/orange-review/app/review/biz/dal/redis"
)

// dal.Init 初始化数据库连接
func Init() {
	// redis.Init()
	mysql.Init()
	es.Init()
}
