package dal

import (
	"github.com/yzc/orange-review/app/id_generator/biz/dal/mysql"
	"github.com/yzc/orange-review/app/id_generator/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
