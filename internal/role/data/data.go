package data

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	perm *mongo.Collection
	role *mongo.Collection
	log  logger.Logger
}

func NewData(db *mongo.Database, name string) *Data {
	return &Data{
		perm: db.Collection("permission"),
		role: db.Collection(name),
		log:  zap.L().Named(name),
	}
}
