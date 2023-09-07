package provider

import (
	"context"
	"ekube/api/pb/token/v1"
	"fmt"
)

var (
	// m is a map from scheme to issuer.
	m = make(map[token.GRANT_TYPE]Issuer)
)

type Issuer interface {
	Init() error
	GrantType() token.GRANT_TYPE
	TokenIssuer
}

// TokenIssuer 访问令牌颁发器
type TokenIssuer interface {
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}

func GetTokenIssuer(gt token.GRANT_TYPE) TokenIssuer {
	if v, ok := m[gt]; ok {
		return v
	}

	return nil
}

// Register 注册
func Register(i Issuer) {
	m[i.GrantType()] = i
}

func Init() error {
	for k, v := range m {
		if err := v.Init(); err != nil {
			return fmt.Errorf("init %s issuer error", k)
		}
	}

	return nil
}
