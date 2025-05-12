package config

type Config struct {
	MySQL   Mysql
	Redis   Redis
	Jwt     Jwt
	Loggger Loggger
	Server  Server
}
