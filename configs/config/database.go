package config

type Database struct {
	Mysql mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql" description:"数据库"`
}

type mysql struct {
	Driver         string `mapstructure:"driver" json:"driver" yaml:"driver" description:"数据库驱动类型"`
	Username       string `mapstructure:"username" json:"username" yaml:"username" description:"账号"`
	Password       string `mapstructure:"password" json:"password" yaml:"password" description:"密码"`
	Host           string `mapstructure:"host" json:"host" yaml:"host" description:"IP"`
	Port           int    `mapstructure:"port" json:"port" yaml:"port" description:"端口"`
	Dbname         string `mapstructure:"dbname" json:"dbname" yaml:"dbname" description:"数据库名称"`
	TablePrefix    string `mapstructure:"tablePrefix" json:"table_prefix" yaml:"tablePrefix" description:"表前缀"`
	Extra          string `mapstructure:"extra" json:"extra" yaml:"extra" description:"DNS额外参数"`
	MaxIdConnect   int    `mapstructure:"maxIdConnect" json:"max_id_connect" yaml:"maxIdConnect" description:"最大连接数"`
	MaxOpenConnect int    `mapstructure:"maxOpenConnect" json:"max_open_connect" yaml:"maxOpenConnect" description:"最大打开数"`
	Charset        string `mapstructure:"charset" json:"charset" yaml:"charset" description:"字符集"`
}
