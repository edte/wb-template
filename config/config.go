// @program:     web-template
// @file:        config.go
// @author:      edte
// @create:      2020-12-19 20:14
// @description: 
package config

// conf 用于获取默认配置
var (
	conf = Config{}
)

const (
	configFileName = "conf.yml"
)

// Config 配置总 struct
type Config struct {
	Mysql *Mysql    `yaml:"mysql"`
	Redis *Redis    `yaml:"redis"`
	MQ    *RabbitMQ `yaml:"mq"`
	Log   *Log      `yaml:"log"`
}

// Mysql 配置详细信息
type Mysql struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

// MysqlConf 获取 Mysql 配置信息
func MysqlConf() *Mysql {
	return conf.Mysql
}

// Redis 缓存的配置信息
type Redis struct {
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
}

// RedisConf 获取 Redis 配置信息
func RedisConf() *Redis {
	return conf.Redis
}

// RabbitMQ 的配置信息
type RabbitMQ struct {
	URL          string `yaml:"url"`
	ExChangeName string `yaml:"exChangeName"`
}

// MQConf 获取 RabbitMQ 详细信息
func MQConf() *RabbitMQ {
	return conf.MQ
}

// Log 的配置信息
type Log struct {
	Path string `yaml:"path"`
}

// LogConf 获取 Log 的配置信息
func LogConf() *Log {
	return conf.Log
}

// 远程配置
func loadRemoteConfig() {

}

// 本地配置文件
func loadLocationConfig() {

}

// 直接在代码中的一个默认配置
func loadDefaultConfig() {

}
