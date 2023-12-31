package impl_test

import (
	"context"
	"ekube/config"
	"ekube/internal/cluster"
	"ekube/protocol/ioc"
	"github.com/emicklei/go-restful/v3"

	"github.com/infraboard/mcube/logger/zap"
)

var (
	impl cluster.Service
	ctx  = context.Background()
)

func init() {
	DevelopmentSetup()
	impl = ioc.GetInternalApp(cluster.AppName).(cluster.Service)
}

func DevelopmentSetup() {
	// 初始化日志实例
	zap.DevelopmentSetup()

	// 针对http handler的测试需要提前设置默认数据格式
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)

	// 初始化配置, 提前配置好/etc/unit_test.env
	err := config.LoadConfigFromToml("../../../etc/config.toml")
	if err != nil {
		panic(err)
	}

	// 初始化全局app
	if err = ioc.InitAllApp(); err != nil {
		panic(err)
	}
}
