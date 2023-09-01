package cluster

import (
	"context"
	clusterv1 "ekube/api/pb/cluster/v1"
)

const (
	AppName = "k8s_clusters"
)

type Service interface {
	ServiceCluster
}

type ServiceCluster interface {
	clusterv1.RPCServer
	CreateCluster(context.Context, *clusterv1.CreateClusterRequest) (*clusterv1.Cluster, error)
	UpdateCluster(context.Context, *clusterv1.UpdateClusterRequest) (*clusterv1.Cluster, error)
	DeleteCluster(context.Context, *clusterv1.DeleteClusterRequest) (*clusterv1.Cluster, error)
}
