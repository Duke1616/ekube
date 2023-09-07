package token

import (
	"context"
	"ekube/api/pb/token/v1"
)

const (
	AppName = "token"
)

type Service interface {
	token.RPCServer
	// IssueToken 颁发Token
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}
