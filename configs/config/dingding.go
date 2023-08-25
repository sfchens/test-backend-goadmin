package config

type DingDing struct {
	Env         string `mapstructure:"env" json:"env" yaml:"env" description:"accessToken"`
	Url         string `mapstructure:"url" json:"url" yaml:"url" description:"地址"`
	Secret      string `mapstructure:"secret" json:"secret" yaml:"secret" description:"secret"`
	AccessToken string `mapstructure:"accessToken" json:"access_token" yaml:"accessToken" description:"accessToken"`
}
