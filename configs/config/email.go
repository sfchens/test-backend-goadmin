package config

type Email struct {
	Disabled bool   `mapstructure:"disabled" json:"disabled" yaml:"disabled" description:"开关"`
	From     string `mapstructure:"from" json:"from" yaml:"from" description:"发送人"`
	Host     string `mapstructure:"host" json:"host" yaml:"host" description:"服务器"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port" description:"端口"`
	Username string `mapstructure:"username" json:"username" yaml:"username" description:"账号"`
	Password string `mapstructure:"password" json:"password" yaml:"password" description:"密码"`
}
