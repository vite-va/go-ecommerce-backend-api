package initialize

import (
	"github.com/vite-va/go-ecommerce-backend-api/global"
	"github.com/vite-va/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = *logger.NewLogger(global.Config.Logger)
}
