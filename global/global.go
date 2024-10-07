package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/vite-va/go-ecommerce-backend-api/pkg/logger"
	"github.com/vite-va/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)

/*
Config
Redis
Mysql
...
*/
