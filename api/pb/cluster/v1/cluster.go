package v1

import (
	"github.com/imdario/mergo"
	"time"
)

func (x *Cluster) Desensitization() {
	if x.Spec.KubeConfig != "" {
		x.Spec.KubeConfig = "****"
	}
}

func (x *Cluster) Update(req *UpdateClusterRequest) {
	m := x.Meta
	m.UpdateAt = time.Now().Unix()
	m.UpdateBy = req.UpdateBy
	x.Spec = req.Spec
}

func (x *Cluster) Patch(req *UpdateClusterRequest) error {
	m := x.Meta
	m.UpdateAt = time.Now().Unix()
	m.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(x.Spec, req.Spec)
}

func (x *ClusterSet) Add(item *Cluster) {
	x.Items = append(x.Items, item)
}
