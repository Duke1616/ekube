package user

import (
	v1 "ekube/api/pb/user/v1"
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
