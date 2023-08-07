package cluster

import (
	"context"
	v1 "ekube/api/cluster"
)

const (
	AppName = "k8s_clusters"
)

type Service interface {
	ClusterService
}

type ClusterService interface {
	v1.RPCServer
	CreateCluster(context.Context, *v1.CreateClusterRequest) (*v1.Cluster, error)
	UpdateCluster(context.Context, *v1.UpdateClusterRequest) (*v1.Cluster, error)
	DeleteCluster(context.Context, *v1.DeleteClusterRequest) (*v1.Cluster, error)
}
