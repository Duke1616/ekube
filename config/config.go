package config

import (
	"context"
	"ekube/pkg/terminal"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/cache/memory"
	"github.com/infraboard/mcube/cache/redis"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	// singleton instance of config package
	_config = defaultConfig()
)

const (
	// DefaultConfigurationName is the default name of configuration
	defaultConfigurationName = "config"

	// DefaultConfigurationPath the default location of the configuration file
	defaultConfigurationPath = "/etc"

	// CIPHER_TEXT_PREFIX the default location of the configuration file
	CIPHER_TEXT_PREFIX = "@ciphered@"
)

func defaultConfig() *config {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigType("toml")
	viper.SetConfigName(defaultConfigurationName)
	viper.AddConfigPath(path + defaultConfigurationPath)

	// Load from current working directory, only used for debugging
	viper.AddConfigPath(".")

	// Load from Environment variables
	viper.SetEnvPrefix("ekube")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &config{
		cfg:         newConfig(),
		cfgChangeCh: make(chan Config),
		watchOnce:   sync.Once{},
		loadOnce:    sync.Once{},
	}
}

func newConfig() *Config {
	return &Config{
		App:        newDefaultAPP(),
		Log:        newDefaultLog(),
		Mongo:      newDefaultMongoDB(),
		Cache:      newDefaultCache(),
		Etcd:       newDefaultEtcd(),
		Image:      newDefaultImage(),
		Kubernetes: &Kubernetes{},
	}
}

// InitGlobal 注入全局变量
func (c *Config) InitGlobal() error {
	// 加载全局缓存
	if err := c.Cache.LoadCache(); err != nil {
		return fmt.Errorf("load cache error, %s", err)
	}

	return nil
}

func (c *Config) Shutdown(ctx context.Context) {
	// 关闭RPC客户端
	//rpc.C().Close()

	// 关闭数据库连接
	c.Mongo.Close(ctx)
}

type config struct {
	cfg         *Config
	cfgChangeCh chan Config
	watchOnce   sync.Once
	loadOnce    sync.Once
}

func (c *config) watchConfig() <-chan Config {
	c.watchOnce.Do(func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			cfg := newConfig()
			if err := viper.Unmarshal(cfg); err != nil {
				klog.Warningf("config reload error: %v", err)
			} else {
				c.cfgChangeCh <- *cfg
				conf = cfg
				conf.InitGlobal()
			}
		})
	})
	return c.cfgChangeCh
}

// TryLoadFromDisk loads configuration from default location after server startup
// return nil error if configuration file not exists
func TryLoadFromDisk() (*Config, error) {
	return _config.loadFromDisk()
}

// WatchConfigChange return config change channel
func WatchConfigChange() <-chan Config {
	return _config.watchConfig()
}

func (c *config) loadFromDisk() (*Config, error) {
	var err error
	c.loadOnce.Do(func() {
		if err = viper.ReadInConfig(); err != nil {
			return
		}

		if err = viper.Unmarshal(c.cfg); err != nil {
			return
		}
	})

	return c.cfg, err
}

// Config 应用配置
type Config struct {
	App            *app             `toml:"app" mapstructure:"app"`
	Log            *log             `toml:"log" mapstructure:"log"`
	Etcd           *etcd            `toml:"etcd" mapstructure:"etcd"`
	Mongo          *mongodb         `toml:"mongodb" mapstructure:"mongodb"`
	Cache          *_cache          `toml:"cache" mapstructure:"cache"`
	Image          *image           `toml:"image" mapstructure:"image"`
	Kubernetes     *Kubernetes      `toml:"Kubernetes" mapstructure:"Kubernetes"`
	TerminalOption *terminal.Option `toml:"terminal" mapstructure:"terminal"`
}

type app struct {
	Name       string `toml:"name" env:"APP_NAME" mapstructure:"name"`
	EncryptKey string `toml:"encrypt_key" env:"APP_ENCRYPT_KEY" mapstructure:"encrypt_key"`
	HTTP       *http  `toml:"http"`
	GRPC       *grpc  `toml:"grpc"`
}

func (a *app) HTTPPrefix() string {
	return fmt.Sprintf("/%s/api", a.Name)
}

func newDefaultAPP() *app {
	return &app{
		Name:       "ekube",
		EncryptKey: "defualt app encrypt key",
		HTTP:       newDefaultHTTP(),
		GRPC:       newDefaultGRPC(),
	}
}

type http struct {
	Host      string `toml:"host" env:"HTTP_HOST" mapstructure:"host"`
	Port      string `toml:"port" env:"HTTP_PORT" mapstructure:"port"`
	EnableSSL bool   `toml:"enable_ssl" env:"HTTP_ENABLE_SSL" mapstructure:"enable_ssl"`
	CertFile  string `toml:"cert_file" env:"HTTP_CERT_FILE" mapstructure:"cert_file"`
	KeyFile   string `toml:"key_file" env:"HTTP_KEY_FILE" mapstructure:"key_file"`
}

func (a *http) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultHTTP() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8080",
	}
}

type grpc struct {
	Host      string `toml:"host" env:"GRPC_HOST" mapstructure:"host"`
	Port      string `toml:"port" env:"GRPC_PORT" mapstructure:"port"`
	EnableSSL bool   `toml:"enable_ssl" env:"GRPC_ENABLE_SSL" mapstructure:"enable_ssl"`
	CertFile  string `toml:"cert_file" env:"GRPC_CERT_FILE" mapstructure:"cert_file"`
	KeyFile   string `toml:"key_file" env:"GRPC_KEY_FILE" mapstructure:"key_file"`
}

func (a *grpc) Addr() string {
	return a.Host + ":" + a.Port
}

func newDefaultGRPC() *grpc {
	return &grpc{
		Host: "127.0.0.1",
		Port: "18080",
	}
}

type log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL" mapstructure:"level"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR" mapstructure:"path_dir"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT" mapstructure:"format"`
	To      LogTo     `toml:"to" env:"LOG_TO" mapstructure:"to"`
}

func newDefaultLog() *log {
	return &log{
		Level:   "debug",
		PathDir: "logs",
		Format:  "text",
		To:      "stdout",
	}
}

func newDefaultEtcd() *etcd {
	m := &etcd{
		UserName:  "",
		Password:  "",
		Endpoints: []string{"127.0.0.1:2379"},
	}
	return m
}

type etcd struct {
	Endpoints []string `toml:"endpoints" env:"ETCD_ENDPOINTS" envSeparator:"," mapstructure:"endpoints"`
	UserName  string   `toml:"username" env:"ETCD_USERNAME" mapstructure:"username"`
	Password  string   `toml:"password" env:"ETCD_PASSWORD" mapstructure:"password"`

	client *clientv3.Client
	lock   sync.Mutex
}

// Client 获取一个全局的mongodb客户端连接
func (e *etcd) Client() (*clientv3.Client, error) {
	// 加载全局数据量单例
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.client == nil {
		conn, err := e.getClient()
		if err != nil {
			return nil, err
		}
		e.client = conn
	}

	return e.client, nil
}

func (e *etcd) getClient() (*clientv3.Client, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   e.Endpoints,
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		return nil, err
	}

	return etcdClient, nil
}

func newDefaultMongoDB() *mongodb {
	m := &mongodb{
		UserName:       "ekube",
		Password:       "123456",
		Database:       "ekube",
		AuthDB:         "",
		Endpoints:      []string{"127.0.0.1:27017"},
		K8sServiceName: "MONGODB",
	}
	m.LoadK8sEnv()
	return m
}

type mongodb struct {
	Endpoints      []string `toml:"endpoints" env:"MONGO_ENDPOINTS" envSeparator:","  mapstructure:"endpoints"`
	UserName       string   `toml:"username" env:"MONGO_USERNAME" mapstructure:"username"`
	Password       string   `toml:"password" env:"MONGO_PASSWORD" mapstructure:"password"`
	Database       string   `toml:"database" env:"MONGO_DATABASE" mapstructure:"database"`
	AuthDB         string   `toml:"auth_db" env:"MONGO_AUTH_DB" mapstructure:"auth_db"`
	K8sServiceName string   `toml:"k8s_service_name" env:"K8S_SERVICE_NAME" mapstructure:"k8s_service_name"`

	client *mongo.Client
	lock   sync.Mutex
}

func (m *mongodb) GetAuthDB() string {
	if m.AuthDB != "" {
		return m.AuthDB
	}

	return m.Database
}

// 当 Pod 运行在 Node 上，kubelet 会为每个活跃的 Service 添加一组环境变量。
// kubelet 为 Pod 添加环境变量 {SVCNAME}_SERVICE_HOST 和 {SVCNAME}_SERVICE_PORT。
// 这里 Service 的名称需大写，横线被转换成下划线
// 具体请参考: https://kubernetes.io/zh-cn/docs/concepts/services-networking/service/#environment-variables
func (m *mongodb) LoadK8sEnv() {
	host := os.Getenv(fmt.Sprintf("%s_SERVICE_HOST", m.K8sServiceName))
	port := os.Getenv(fmt.Sprintf("%s_SERVICE_PORT", m.K8sServiceName))
	addr := fmt.Sprintf("%s:%s", host, port)
	if host != "" && port != "" {
		m.Endpoints = []string{addr}
	}
}

func (m *mongodb) GetDB() (*mongo.Database, error) {
	conn, err := m.Client()
	if err != nil {
		return nil, err
	}
	return conn.Database(m.Database), nil
}

// 关闭数据库连接
func (m *mongodb) Close(ctx context.Context) error {
	if m.client == nil {
		return nil
	}

	return m.client.Disconnect(ctx)
}

// Client 获取一个全局的mongodb客户端连接
func (m *mongodb) Client() (*mongo.Client, error) {
	// 加载全局数据量单例
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.client == nil {
		conn, err := m.getClient()
		if err != nil {
			return nil, err
		}
		m.client = conn
	}

	return m.client, nil
}

func (m *mongodb) getClient() (*mongo.Client, error) {
	opts := options.Client()

	if m.UserName != "" && m.Password != "" {
		cred := options.Credential{
			AuthSource: m.GetAuthDB(),
		}
		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true
		opts.SetAuth(cred)
	}
	opts.SetHosts(m.Endpoints)
	opts.SetConnectTimeout(5 * time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("new mongodb client error, %s", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("ping mongodb server(%s) error, %s", m.Endpoints, err)
	}

	return client, nil
}

func newDefaultImage() *image {
	return &image{
		DefaultRegistry: "registry.cn-hangzhou.aliyuncs.com",
	}
}

type image struct {
	// 镜像默认推送仓库地址
	DefaultRegistry string `toml:"default_registry" json:"default_registry" yaml:"default_registry" env:"DEFAULT_REGISTRY"`
}

func newDefaultCache() *_cache {
	return &_cache{
		Type:   "memory",
		Memory: memory.NewDefaultConfig(),
		Redis:  redis.NewDefaultConfig(),
	}
}

type _cache struct {
	Type   string         `toml:"type" json:"type" yaml:"type" env:"MCENTER_CACHE_TYPE" mapstructure:"type"`
	Memory *memory.Config `toml:"memory" json:"memory" yaml:"memory" mapstructure:"memory"`
	Redis  *redis.Config  `toml:"redis" json:"redis" yaml:"redis" mapstructure:"redis"`
}

func (c *_cache) LoadCache() error {
	// 设置全局缓存
	switch c.Type {
	case "memory", "":
		ins := memory.NewCache(c.Memory)
		if err := ins.ClearAll(); err != nil {
			return err
		}

		cache.SetGlobal(ins)
		zap.L().Info("use cache in local memory")
	case "redis":
		ins := redis.NewCache(c.Redis)
		cache.SetGlobal(ins)
		zap.L().Info("use redis to cache")
	default:
		return fmt.Errorf("unknown cache type: %s", c.Type)
	}

	return nil
}

type Kubernetes struct {
	// kubeconfig path, if not specified, will use
	// in cluster way to create clientset
	KubeConfig string `json:"kubeconfig" yaml:"kubeconfig"`

	// Kubernetes apiserver public address, used to generate kubeconfig
	// for downloading, default to host defined in kubeconfig
	// +optional
	Master string `json:"master,omitempty" yaml:"master,omitempty"`

	// Kubernetes clientset qps
	// +optional
	QPS float32 `json:"qps,omitempty" yaml:"qps,omitempty"`

	// Kubernetes clientset burst
	// +optional
	Burst int `json:"burst,omitempty" yaml:"burst,omitempty"`
}

func (k *Kubernetes) Validate() []error {
	var errors []error

	if k.KubeConfig != "" {
		if _, err := os.Stat(k.KubeConfig); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
