package sys_model

type SysConfig struct {
	ID        uint        `json:"id" description:"ID"`
	Name      string      ` json:"name" description:"名称"`
	Type      string      ` json:"type" description:"类型"`
	Config    interface{} `json:"config" description:"配置"`
	Operator  string      `json:"operator" description:"操作人"`
	CreatedAt string      `json:"created_at" description:"创建时间"`
	UpdatedAt string      `json:"updated_at" description:"更新时间"`
}

type SysConfigJson struct {
	SiteName       string `json:"site_name" description:"站点名称"`
	SiteUrl        string `json:"site_url" description:"站点名称"`
	SiteLogo       string `json:"site_logo" description:"站点名称"`
	TengxunMapKey  string `json:"tengxun_map_key" description:"站点名称"`
	SiteLogoSquare string `json:"site_logo_square" description:"站点名称"`
	LoginLogo      string `json:"login_logo" description:"站点名称"`
	WapLoginLogo   string `json:"wap_login_logo" description:"站点名称"`
	StationOpen    string `json:"station_open" description:"站点名称"`
}
