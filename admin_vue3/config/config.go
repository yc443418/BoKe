// 读取配置文件
// @author chen
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// config 总配置文件
type config struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
	Mysql  mysql  `yaml:"mysql"`
	Redis  redis  `yaml:"redis"`
	Token  token  `yaml:"token"`
}

// 系统配置
type system struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

// logger配置
type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"showLine"`
	LogInConsole bool   `yaml:"logInConsole"`
}

// mysql 配置
type mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"logLevel"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// token配置
type token struct {
	Headers    string `yaml:"headers"`
	ExpireTime int    `yaml:"expireTime"`
	Secret     string `yaml:"secret"`
	Issuer     string `yaml:"issuer"`
}

var Config *config

// init 初始化配置
func init() {

	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	yaml.Unmarshal(yamlFile, &Config)
}
