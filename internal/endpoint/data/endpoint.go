package data

import (
	"context"
	v1 "ekube/api/pb/endpoint/v1"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Data) InsertMany(ctx context.Context, endpoints []*v1.Endpoint) error {
	news := make([]interface{}, 0, len(endpoints))

	// 检测是否已经注册
	for i := range endpoints {
		if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": endpoints[i].Id}, endpoints[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, endpoints[i])
			} else {
				return err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(ctx, news); err != nil {
			return exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}
	return nil
}

func (s *Data) Get(ctx context.Context, req *v1.DescribeEndpointRequest) (*v1.Endpoint, error) {
	filter := bson.M{}
	if req.Id != "" {
		filter["_id"] = req.Id
	}

	resp := &v1.Endpoint{}

	if err := s.col.FindOne(ctx, filter).Decode(resp); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("endpoint %s not found", req)
		}

		return nil, exception.NewInternalServerError("find endpoint %s error, %s", req.Id, err)
	}

	return resp, nil
}
