package Config

/*
远程连接属性
*/
type BasicHost struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Key      string `yaml:"Key"`
	Password string `yaml:"Password"`
}

/*
鉴权对象
*/
type AuthConfig struct {
	NeedAuth bool     `yaml:"NeedAuth"`
	List     []string `yaml:"List"`
}

/*
数据库连接配置
*/
type DBConfig struct {
	CurrentSource string   `yaml:"CurrentSource"`
	DBInfo        []DBConn `yaml:"DBInfo"`
}

/*
数据库连接配置属性详情
*/
type DBConn struct {
	DBName           string `yaml:"DBName"`
	User             string `yaml:"User"`
	Password         string `yaml:"Password"`
	Host             string `yaml:"Host"`
	Driver           string `yaml:"Driver"`
	DefaultCharset   string `yaml:"DefaultCharset"`
	DefaultCollation string `yaml:"DefaultCollation"`
	MaxConn          int    `yaml:"MaxConn"`
	MaxAlive         int    `yaml:"MaxAlive"`
	ParseTime        bool   `yaml:"ParseTime"`
	Location         string `yaml:"Location"`
	Port             int    `yaml:"Port"`
	Key              string `yaml:"Key"`
}

/*
基础配置
*/
type PropsConfig struct {
	TestSite   bool        `yaml:"TestSite"`
	ServerPort string      `yaml:"ServerPort"`
	Mail       []BasicHost `yaml:"Mail"`
	FTP        []BasicHost `yaml:"FTP"`
}
