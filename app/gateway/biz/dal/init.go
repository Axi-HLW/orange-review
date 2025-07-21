package dal

import (
	"github.com/yzc/orange-review/app/gateway/biz/dal/mysql"
	"github.com/yzc/orange-review/app/gateway/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
