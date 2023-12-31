package cmd

import (
	"ekube/cmd/controller"
	option1 "ekube/cmd/controller/option"
	"ekube/cmd/start"
	"ekube/cmd/start/option"
	"ekube/config"
	"ekube/protocol/ioc"
	"ekube/version"
	"errors"
	"fmt"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	"github.com/infraboard/mcube/validator"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mpaas",
	Short: "微服务发布平台",
	Long:  "微服务发布平台",
	Run: func(cmd *cobra.Command, args []string) {
		if vers {
			fmt.Println(version.FullVersion())
			return
		}
		cmd.Help()
	},
}

// config 为全局变量, 只需要load 即可全局可用户
func loadGlobalConfig(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := config.LoadConfigFromToml()
		if err != nil {
			return err
		}
	case "env":
		err := config.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown config type")
	}

	return nil
}

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	lc := config.C().Log
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level
	switch lc.To {
	case config.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case config.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	switch lc.Format {
	case config.JSONFormat:
		zapConfig.JSON = true
	}
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("init").Info(logInitMsg)
	return nil
}

func initail() {
	err := validator.Init()
	cobra.CheckErr(err)

	// 初始化全局变量
	err = loadGlobalConfig(confType)
	cobra.CheckErr(err)

	// 初始化全局日志配置
	err = loadGlobalLogger()
	cobra.CheckErr(err)

	// 初始化全局app
	err = ioc.InitAllApp()
	cobra.CheckErr(err)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// 补充初始化设置
	cobra.OnInitialize(initail)
	s := option.NewServerRunOptions()
	r := option1.NewKubeSphereControllerManagerOptions()

	apiServerCmd := start.NewAPIServerCommand(s)
	controllerCmd := controller.NewControllerCommand(r)
	RootCmd.AddCommand(apiServerCmd)
	RootCmd.AddCommand(controllerCmd)
	//RootCmd.AddCommand(initial.Cmd)

	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")
	RootCmd.PersistentFlags().StringVarP(&confETCD, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the mpaas version")
}
