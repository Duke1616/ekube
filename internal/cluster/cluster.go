package cluster

import (
	clusterv1 "ekube/api/pb/cluster/v1"
	"ekube/tools"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"time"

	pb_request "github.com/infraboard/mcube/pb/request"
)

func NewCluster(req *clusterv1.CreateClusterRequest) (*clusterv1.Cluster, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &clusterv1.Cluster{
		Meta:   tools.NewMeta(),
		Spec:   req,
		Info:   &clusterv1.ServerInfo{},
		Status: &clusterv1.Status{},
	}, nil
}

func NewCreateClusterRequest() *clusterv1.CreateClusterRequest {
	return &clusterv1.CreateClusterRequest{}
}

func NewDefaultCluster() *clusterv1.Cluster {
	return &clusterv1.Cluster{
		Spec: &clusterv1.CreateClusterRequest{},
	}
}

func NewDescribeClusterRequest(id string) *clusterv1.DescribeClusterRequest {
	return &clusterv1.DescribeClusterRequest{
		Id: id,
	}
}

func NewDeleteClusterRequestWithID(id string) *clusterv1.DeleteClusterRequest {
	return &clusterv1.DeleteClusterRequest{
		Id: id,
	}
}

func NewPutClusterRequest(id string) *clusterv1.UpdateClusterRequest {
	return &clusterv1.UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMicro(),
		Spec:       NewCreateClusterRequest(),
	}
}

func NewPatchClusterRequest(id string) *clusterv1.UpdateClusterRequest {
	return &clusterv1.UpdateClusterRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMicro(),
		Spec:       NewCreateClusterRequest(),
	}
}

func NewClusterSet() *clusterv1.ClusterSet {
	return &clusterv1.ClusterSet{
		Items: []*clusterv1.Cluster{},
	}
}

func NewListClusterRequestFromHTTP(r *restful.Request) *clusterv1.ListClusterRequest {
	req := NewListClusterRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Keywords = r.QueryParameter("keywords")
	req.Vendor = r.QueryParameter("vendor")
	req.Region = r.QueryParameter("region")
	return req
}

func NewListClusterRequest() *clusterv1.ListClusterRequest {
	return &clusterv1.ListClusterRequest{
		Page: request.NewDefaultPageRequest(),
	}
}
