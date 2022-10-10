package global

import (
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/config"
)

var (
	CONFIG config.Config
	VP     *viper.Viper
	DB     *gorm.DB
	REDIS  *redis.Client
	LOG    *zap.SugaredLogger

	//
	//
	// Timer               timer.Timer = timer.NewTimerTask()
	// GVA_Concurrency_Control             = &singleflight.Group{}
	// BlackCache local_cache.Cache
	// lock       sync.RWMutex
)
