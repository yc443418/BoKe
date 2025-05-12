package global

import (
	"Traceability/config"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

var (
	CONFIG config.Config
	DB     *gorm.DB
	Cron   *cron.Cron
	Redis  *redis.Client
	Logger *zap.Logger
)
