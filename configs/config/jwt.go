package config

type Jwt struct {
	Secret      string `mapstructure:"secret" json:"secret" yaml:"secret" description:"密钥"`
	ExpiresTime int    `mapstructure:"expiresTime" json:"expires_time" yaml:"expiresTime" description:"过期时间"`
	SigningKey  string `mapstructure:"signingKey" json:"signing_key" yaml:"signingKey" description:"端口"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer" description:"读超时时间"`
}
