package dal

import (
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/auth/biz/dal/redis"
)

func InitRepo() {
	redis.Init()
	mysql.Init()
}
