package initialize

import (
	"context"
	"go.uber.org/zap"
	"server/global"
	"server/model/entity"
)

func InitBlackList() {
	var data []string
	err := global.DB.Model(&entity.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		//global.BlackCache.SetDefault(data[i], struct{}{})
		global.REDIS.Set(context.Background(), data[i], struct{}{}, 0)
	} // jwt黑名单 加入 BlackCache 中
}
