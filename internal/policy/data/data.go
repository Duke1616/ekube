package data

import (
	"ekube/internal/role"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	col  *mongo.Collection
	role role.Service
	log  logger.Logger
}

func NewData(db *mongo.Database, name string) *Data {
	return &Data{
		col:  db.Collection(name),
		log:  zap.L().Named(name),
		role: ioc.GetInternalApp(role.AppName).(role.Service),
	}
}
