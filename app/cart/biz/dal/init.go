package dal

import (
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/dal/mysql"
	"github.com/Cui-Guo-crushed-his-team/CuiGuoMall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
