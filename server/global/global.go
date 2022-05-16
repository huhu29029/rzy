package global

import (
	"go.uber.org/zap"

	"online-exam/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	MongoDB *mongo.Client
	REDIS   *redis.Client
	CONFIG  config.Server
	VP      *viper.Viper
	//LOG    *oplogging.Logger
	LOG *zap.Logger
)
