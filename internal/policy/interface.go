package policy

import (
	"context"
	"ekube/api/pb/policy/v1"
)

const (
	AppName = "policy"
)

type Service interface {
	policy.RPCServer
	CreatePolicy(context.Context, *policy.CreatePolicyRequest) (*policy.Policy, error)
}
