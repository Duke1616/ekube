package deployment

import (
	"context"
	"ekube/pkg/k8s/meta"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetDeployment(ctx context.Context, req *meta.GetRequest) (*appsv1.Deployment, error) {
	d, err := c.appsv1.Deployments(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	d.APIVersion = "apps/v1"
	d.Kind = "Deployment"
	return d, nil
}

func (c *Client) ListDeployment(ctx context.Context, req *meta.ListRequest) (*appsv1.DeploymentList, error) {
	ds, err := c.appsv1.Deployments(req.Namespace).List(ctx, req.Opts)
	if err != nil {
		return nil, err
	}
	if req.SkipManagedFields {
		for i := range ds.Items {
			ds.Items[i].ManagedFields = nil
		}
	}
	return ds, nil
}
