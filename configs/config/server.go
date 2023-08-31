package config

type Server struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	DingDing DingDing `mapstructure:"dingDing" json:"ding_ding" yaml:"dingDing"`
	Email    Email    `mapstructure:"email" json:"email" yaml:"email"`
}
