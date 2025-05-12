package config

// JWT 配置结构体
type Jwt struct {
	Issuer     string `json:"issuer"`
	Secret     string `json:"secret"`
	ExpireTime int64  `json:"expireTime"`
	NotBefore  int64  `json:"notBefore"`
}
