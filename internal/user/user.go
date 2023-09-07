package user

import (
	v1 "ekube/api/pb/user/v1"
	"ekube/tools"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// NewHashedPassword 生产hash后的密码对象
func NewHashedPassword(password string) (*v1.Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return &v1.Password{
		Password:      string(bytes),
		CreateAt:      time.Now().UnixMilli(),
		UpdateAt:      time.Now().UnixMilli(),
		ExpiredDays:   90,
		ExpiredRemind: 30,
	}, nil
}

func NewDescribeUserRequestByName(username string) *v1.DescribeUserRequest {
	return &v1.DescribeUserRequest{
		Username: username,
	}
}

func TypeToString(uts ...v1.TYPE) (types []string) {
	for _, t := range uts {
		types = append(types, t.String())
	}
	return
}

func NewCreateUserRequest() *v1.CreateUserRequest {
	return &v1.CreateUserRequest{
		Labels:  map[string]string{},
		Profile: &v1.Profile{},
	}
}

func NewUserSet() *v1.UserSet {
	return &v1.UserSet{
		Items: []*v1.User{},
	}
}

func NewListUserRequestFromHTTP(r *http.Request) *v1.ListUserRequest {
	query := NewListUserRequest()

	qs := r.URL.Query()
	query.Page = tools.NewPageRequestFromHTTP(r)
	//query.Keywords = qs.Get("keywords")
	//query.SkipItems = qs.Get("skip_items") == "true"

	uids := qs.Get("user_ids")
	if uids != "" {
		query.UserIds = strings.Split(uids, ",")
	}
	return query
}

// NewQueryUserRequest 列表查询请求
func NewListUserRequest() *v1.ListUserRequest {
	return &v1.ListUserRequest{
		Page: tools.NewPageRequest(20, 1),
	}
}
