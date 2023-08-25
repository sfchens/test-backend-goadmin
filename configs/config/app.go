package config

type App struct {
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode" description:"开环境"`
	Name          string `mapstructure:"name" json:"name" yaml:"name" description:"站点名称"`
	Port          int    `mapstructure:"port" json:"port" yaml:"port" description:"端口"`
	ReadTimeout   int    `mapstructure:"readTimeout" json:"read_timeout" yaml:"readTimeout" description:"读超时时间"`
	WriterTimeout int    `mapstructure:"writerTimeout" json:"writer_timeout" yaml:"writerTimeout" description:"写超时时间"`
	BaseUrl       string `mapstructure:"baseUrl" json:"base_url" yaml:"baseUrl" description:"站点地址"`
	IsApiMysql    bool   `mapstructure:"isApiMysql" json:"is_api_mysql" yaml:"isApiMysql" description:"站点地址"`
}
