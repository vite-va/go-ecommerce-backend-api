package initialize

import (
	"github.com/vite-va/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	// load configuration
	LoadConfig()
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))
	InitMysql()
	// InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
