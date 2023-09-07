package token

import (
	v1 "ekube/api/pb/token/v1"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger/zap"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	// token默认过期时长
	DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND = 3600
	// 刷新token默认过期时间
	DEFAULT_REFRESH_TOKEN_EXPIRE_SECOND = DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND * 4
)

const (
	// 携带Token的Header 名称, 标准格式: Authorization: bearer xxxx-token
	ACCESS_TOKEN_HEADER_KEY = "Authorization"
	// token也可以放到Cookie中, 只是cookie本身并不是太安全, 如果Authorization头没有才考虑从
	// cookie中获取
	ACCESS_TOKEN_COOKIE_KEY = "ekube.access_token"
	// token也可以放到Query参数中, 只是Query参数不安全, 容易导致Token泄露, 一般只用于调试
	ACCESS_TOKEN_QUERY_KEY = "ekube_access_token"
)

func NewToken(req *v1.IssueTokenRequest) *v1.Token {
	tk := &v1.Token{
		AccessToken:      MakeBearer(24),
		RefreshToken:     MakeBearer(32),
		IssueAt:          time.Now().Unix(),
		AccessExpiredAt:  req.ExpiredAt,
		RefreshExpiredAt: req.ExpiredAt * 4,
		GrantType:        req.GrantType,
		Type:             req.Type,
		Description:      req.Description,
		Meta:             map[string]string{},
	}
	switch req.GrantType {
	case v1.GRANT_TYPE_PASSWORD:
		tk.Platform = v1.PLATFORM_API
	default:
		tk.Platform = v1.PLATFORM_WEB
	}
	return tk
}

func MakeBearer(lenth int) string {
	charList := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	t := make([]string, lenth)
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(lenth) + rand.Int63n(10000)))
	for i := 0; i < lenth; i++ {
		rn := r.Intn(len(charList))
		w := charList[rn : rn+1]
		t = append(t, w)
	}

	token := strings.Join(t, "")
	return token
}

func NewIssueTokenRequest() *v1.IssueTokenRequest {
	return &v1.IssueTokenRequest{
		ExpiredAt: DEFAULT_ACCESS_TOKEN_EXPIRE_SECOND,
	}
}

func GetAccessTokenFromHTTP(r *http.Request) string {
	auth := r.Header.Get(ACCESS_TOKEN_HEADER_KEY)
	info := strings.Split(auth, " ")
	if len(info) > 1 {
		return info[1]
	}

	ck, err := r.Cookie(ACCESS_TOKEN_COOKIE_KEY)
	if err != nil {
		zap.L().Warnf("get tk from cookie: %s error, %s", ACCESS_TOKEN_COOKIE_KEY, err)
		return r.URL.Query().Get(ACCESS_TOKEN_QUERY_KEY)
	}

	return ck.Value
}

func NewValidateTokenRequest(accessToken string) *v1.ValidateTokenRequest {
	return &v1.ValidateTokenRequest{
		AccessToken: accessToken,
	}
}

func GetTokenFromRequest(r *restful.Request) *v1.Token {
	tk := r.Attribute("token")
	if tk == nil {
		return nil
	}
	return tk.(*v1.Token)
}
