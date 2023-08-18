package impl

import (
	"context"
	clusterv1 "ekube/api/pb/cluster/v1"
	"ekube/internal/cluster"
	"ekube/internal/cluster/data"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
)

// CreateCluster 集群录入
func (s *service) CreateCluster(ctx context.Context, req *clusterv1.CreateClusterRequest) (
	*clusterv1.Cluster, error) {
	ins, err := cluster.NewCluster(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create cluster error, %s", err)
	}

	// 连接集群检查状态
	s.checkStatus(ins)
	if err = ins.IsAlive(); err != nil {
		return nil, err
	}

	// 加密处理
	err = ins.EncryptKubeConf(s.encryptKey)
	if err != nil {
		return nil, err
	}

	if err = s.data.Save(ctx, ins); err != nil {
		return nil, err
	}

	return ins, err
}

func (s *service) checkStatus(ins *clusterv1.Cluster) {
	//client, err := resource.NewClient(ins.Spec.KubeConfig)
	//if err != nil {
	//	ins.Status.Message = err.Error()
	//	return
	//}
	//
	//if ctx := client.CurrentContext(); ctx != nil {
	//	ins.Meta.Id = ctx.Cluster
	//	ins.Meta.ServerInfo.AuthUser = ctx.AuthInfo
	//}
	//
	//if k := client.CurrentCluster(); k != nil {
	//	ins.Meta.ServerInfo.Server = k.Server
	//}
	//
	//// 检查凭证是否可用
	//ins.Status.CheckAt = time.Now().Unix()
	//v, err := client.ServerVersion()
	//if err != nil {
	//	ins.Status.IsAlive = false
	//	ins.Status.Message = err.Error()
	//} else {
	//	ins.Status.IsAlive = true
	//	ins.Meta.ServerInfo.Version = v
	//}
}

// ListCluster 查询集群列表
func (s *service) ListCluster(ctx context.Context, req *clusterv1.ListClusterRequest) (
	*clusterv1.ClusterSet, error) {
	query := data.NewListClusterRequest(req)
	set, err := s.data.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return set, nil
}

// DescribeCluster 查询集群详情
func (s *service) DescribeCluster(ctx context.Context, req *clusterv1.DescribeClusterRequest) (
	*clusterv1.Cluster, error) {
	ins, err := s.data.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if err = ins.DecryptKubeConf(s.encryptKey); err != nil {
		return nil, err
	}
	return ins, nil
}

// UpdateCluster 集群更新
func (s *service) UpdateCluster(ctx context.Context, req *clusterv1.UpdateClusterRequest) (
	*clusterv1.Cluster, error) {
	ins, err := s.DescribeCluster(ctx, cluster.NewDescribeClusterRequest(req.Id))
	if err != nil {
		return nil, err
	}

	// 配置kubeconfig是否有变更
	isKubeConfigChanged := req.Spec.KubeConfig == ins.Spec.KubeConfig

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err = ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	// 校验更新后数据合法性
	if err = ins.Spec.Validate(); err != nil {
		return nil, err
	}

	// 如果有变更检查集群状态
	if isKubeConfigChanged {
		s.checkStatus(ins)
	}

	// 加密
	err = ins.EncryptKubeConf(s.encryptKey)
	if err != nil {
		return nil, err
	}

	if err = s.data.Update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

// DeleteCluster 集群的删除
func (s *service) DeleteCluster(ctx context.Context, req *clusterv1.DeleteClusterRequest) (
	*clusterv1.Cluster, error) {

	ins, err := s.DescribeCluster(ctx, cluster.NewDescribeClusterRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err = s.data.Delete(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}
