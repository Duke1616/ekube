package config

import (
	"github.com/caarlos0/env/v6"
	"k8s.io/klog/v2"
)

var (
	conf *Config
)

// C 全局配置对象
func C() *Config {
	if conf == nil {
		panic("Load Config first")
	}

	return conf
}

func LoadConfigFromToml(filePath ...string) error {
	var err error
	conf, err = TryLoadFromDisk()
	if err == nil {
		//conf.Kubernetes = s.KubernetesOption
	} else {
		klog.Fatalf("Failed to load configuration from disk: %v", err)
	}

	return conf.InitGlobal()
}

//// LoadConfigFromTomlV1 从toml中添加配置文件, 并初始化全局对象
//func LoadConfigFromToml(filePath string) error {
//	conf = newConfig()
//	if _, err := toml.DecodeFile(filePath, conf); err != nil {
//		return err
//	}
//
//	return conf.InitGlobal()
//}

// LoadConfigFromEnv 从环境变量中加载配置
func LoadConfigFromEnv() error {
	conf = newConfig()
	if err := env.Parse(conf); err != nil {
		return err
	}
	return conf.InitGlobal()
}
