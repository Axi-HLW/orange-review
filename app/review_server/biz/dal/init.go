package dal

import (
	"github.com/yzc/orange-review/app/review_server/biz/dal/mysql"
	"github.com/yzc/orange-review/app/review_server/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
