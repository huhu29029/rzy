package initialize

import (
	"context"
	"online-exam/global"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func MongoDB() *mongo.Client {
	// uri := "mongodb://daima:d8fk2xOvwDoHBiZRuC3X@127.0.0.1:27017/?authSource=daima"
	uri := global.CONFIG.Mongo.Uri
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		global.LOG.Error("mongodb connect failed", zap.Any("err", err))
	}
	return client
}
