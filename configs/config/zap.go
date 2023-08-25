package config

type Zap struct {
	Level         string   `mapstructure:"level" json:"level" yaml:"level" description:"等级"`
	Prefix        string   `mapstructure:"prefix" json:"prefix" yaml:"prefix" description:"前缀"`
	Format        string   `mapstructure:"format" json:"format" yaml:"format" description:"格式"`
	Path          string   `mapstructure:"path" json:"path" yaml:"path" description:"路径"`
	LogFile       []string `mapstructure:"logFile" json:"log_file" yaml:"logFile" description:"文件路径"`
	EncodeLevel   string   `mapstructure:"encodeLevel" json:"encode_level" yaml:"encodeLevel" description:"等级"`
	StacktraceKey string   `mapstructure:"stacktraceKey" json:"stacktrace_key" yaml:"stacktraceKey" description:"等级"`
	MaxAge        int      `mapstructure:"maxAge" json:"max_age" yaml:"maxAge" description:"等级"`
	ShowLine      bool     `mapstructure:"showLine" json:"show_line" yaml:"showLine" description:"等级"`
	LogInConsole  bool     `mapstructure:"logInConsole" json:"log_in_console" yaml:"logInConsole" description:"显示console"`
}
