package password

import (
	"context"
	tokenV1 "ekube/api/pb/token/v1"
	userV1 "ekube/api/pb/user/v1"
	"ekube/internal/token"
	"ekube/internal/token/provider"
	"ekube/internal/user"
	"ekube/protocol/ioc"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	AUTH_FAILED = exception.NewUnauthorized("user or password not connrect")
)

type issuer struct {
	user user.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.user = ioc.GetInternalApp(user.AppName).(user.Service)
	i.log = zap.L().Named("issuer.password")
	return nil
}

func (i *issuer) GrantType() tokenV1.GRANT_TYPE {
	return tokenV1.GRANT_TYPE_PASSWORD
}

func (i *issuer) validate(ctx context.Context, username, pass string) (*userV1.User, error) {
	if username == "" || pass == "" {
		return nil, AUTH_FAILED
	}

	// 检测用户的密码是否正确
	u, err := i.user.DescribeUser(ctx, user.NewDescribeUserRequestByName(username))
	if err != nil {
		return nil, err
	}

	//if err = u.Password.CheckPassword(pass); err != nil {
	//	return nil, AUTH_FAILED
	//}

	return u, nil
}

func (i *issuer) IssueToken(ctx context.Context, req *tokenV1.IssueTokenRequest) (*tokenV1.Token, error) {
	u, err := i.validate(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// 3. 颁发Token
	tk := token.NewToken(req)
	tk.Username = u.Spec.Username
	tk.UserType = u.Spec.Type
	tk.UserId = u.Meta.Id
	return tk, nil
}

func init() {
	provider.Register(&issuer{})
}
