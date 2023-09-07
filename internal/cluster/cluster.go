package cluster

import (
	v1 "ekube/api/pb/cluster/v1"
	"ekube/tools"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"time"

	pb_request "github.com/infraboard/mcube/pb/request"
)

func NewCluster(req *v1.CreateClusterRequest) (*v1.Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &v1.Cluster{
		Meta:   tools.NewMeta(),
		Spec:   req,
		Info:   &v1.ServerInfo{},
		Status: &v1.Status{},
	}, nil
}

func NewCreateClusterRequest() *v1.CreateClusterRequest {
	return &v1.CreateClusterRequest{}
}

func NewDefaultCluster() *v1.Cluster {
	return &v1.Cluster{
		Spec: &v1.CreateClusterRequest{},
	}
}

func NewDescribeClusterRequest(id string) *v1.DescribeClusterRequest {
	return &v1.DescribeClusterRequest{
		Id: id,
	}
}

func NewDeleteClusterRequestWithID(id string) *v1.DeleteClusterRequest {
	return &v1.DeleteClusterRequest{
		Id: id,
	}
}

func NewPutClusterRequest(id string) *v1.UpdateClusterRequest {
	return &v1.UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Spec:       NewCreateClusterRequest(),
	}
}

func NewPatchClusterRequest(id string) *v1.UpdateClusterRequest {
	return &v1.UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Spec:       NewCreateClusterRequest(),
	}
}

func NewClusterSet() *v1.ClusterSet {
	return &v1.ClusterSet{
		Items: []*v1.Cluster{},
	}
}

func NewListClusterRequestFromHTTP(r *restful.Request) *v1.ListClusterRequest {
	req := NewListClusterRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Keywords = r.QueryParameter("keywords")
	req.Vendor = r.QueryParameter("vendor")
	req.Region = r.QueryParameter("region")
	return req
}

func NewListClusterRequest() *v1.ListClusterRequest {
	return &v1.ListClusterRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
