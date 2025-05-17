// 读取配置文件
// @author chen

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// config 总配置文件
type config struct {
	System         system         `yaml:"system"`
	Logger         logger         `yaml:"logger"`
	Mysql          mysql          `yaml:"mysql"`
	Redis          redis          `yaml:"redis"`
	Mail           mail           `yaml:"mail"`
	Token          token          `yaml:"token"`
	UploadSettings uploadSettings `yaml:"uploadSettings"`
}

// system 配置
type system struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

// logger 日志
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
	LogLevel string `yaml:"log_level"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

// redis 配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// QQ邮箱
type mail struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

// token 配置
type token struct {
	Headers    string `yaml:"headers"`
	ExpireTime int    `yaml:"expireTime"`
	Secret     string `yaml:"secret"`
	Issuer     string `yaml:"issuer"`
}

// uploadSettings文件地址和ip
type uploadSettings struct {
	UploadDir  string `yaml:"uploadDir"`
	UploadHost string `yaml:"uploadHost"`
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
