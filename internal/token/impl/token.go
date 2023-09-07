package impl

import (
	"context"
	v1 "ekube/api/pb/token/v1"
	"ekube/internal/token/provider"
	"github.com/infraboard/mcube/exception"
)

func (s *service) IssueToken(ctx context.Context, req *v1.IssueTokenRequest) (*v1.Token, error) {
	// 颁发令牌
	tk, err := s.IssueTokenNow(ctx, req)
	if err != nil {
		return nil, err
	}

	return tk, nil
}

func (s *service) IssueTokenNow(ctx context.Context, req *v1.IssueTokenRequest) (*v1.Token, error) {
	// 获取令牌颁发器
	issuer := provider.GetTokenIssuer(req.GrantType)

	// 确保有provider
	if issuer == nil {
		return nil, exception.NewBadRequest("grant type %s not support", req.GrantType)
	}

	// 颁发token
	tk, err := issuer.IssueToken(ctx, req)
	if err != nil {
		return nil, err
	}

	if !req.DryRun {
		// 入库保存
		if err = s.data.Insert(ctx, tk); err != nil {
			return nil, err
		}

	}

	return tk, nil
}

func (s *service) ValidateToken(ctx context.Context, req *v1.ValidateTokenRequest) (*v1.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.data.Get(ctx, req.AccessToken)
	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	tk.RefreshToken = ""

	return tk, nil
}
