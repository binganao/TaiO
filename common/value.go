package common

import "github.com/spf13/viper"

var (

	/**
	节点类型
	*/
	CONFIG_NODE_TYPE = "node.type"
	CONFIG_DATA_ADDR = "node.addr"
	NODE_TYPE        string
	DATA_ADDR        string

	/**
	数据库配置文件节点
	*/
	CONFIG_DATABASE_DRIVER   = "database.driver"
	CONFIG_DATABASE_HOST     = "database.mysql.host"
	CONFIG_DATABASE_PORT     = "database.mysql.port"
	CONFIG_DATABASE_USERNAME = "database.mysql.username"
	CONFIG_DATABASE_PASSWORD = "database.mysql.password"
	CONFIG_DATABASE_DATABASE = "database.mysql.database"
	CONFIG_DATABASE_CHARSET  = "database.mysql.charset"

	DATABASE_DRIVER   string // 数据库类型
	DATABASE_HOST     string // 数据库 IP
	DATABASE_PORT     string // 数据库端口
	DATABASE_USERNAME string // 数据库用户名
	DATABASE_PASSWORD string // 数据库密码
	DATABASE_DATABASE string // 数据库名称
	DATABASE_CHARSET  string // 数据库 字符集

	/**
	MASSCAN
	*/
	CONFIG_MASSCAN_RATE = "masscan.rate"
	MASSCAN_RATE        int

	/**
	SECRET
	*/
	CONFIG_SECRET = "DATA.SECRET"
	SECRET        string
)

func InitValue() {

	NODE_TYPE = viper.GetString(CONFIG_NODE_TYPE)
	DATA_ADDR = viper.GetString(CONFIG_DATA_ADDR)

	DATABASE_DRIVER = viper.GetString(CONFIG_DATABASE_DRIVER)
	DATABASE_HOST = viper.GetString(CONFIG_DATABASE_HOST)
	DATABASE_PORT = viper.GetString(CONFIG_DATABASE_PORT)
	DATABASE_USERNAME = viper.GetString(CONFIG_DATABASE_USERNAME)
	DATABASE_PASSWORD = viper.GetString(CONFIG_DATABASE_PASSWORD)
	DATABASE_DATABASE = viper.GetString(CONFIG_DATABASE_DATABASE)
	DATABASE_CHARSET = viper.GetString(CONFIG_DATABASE_CHARSET)

	MASSCAN_RATE = viper.GetInt(CONFIG_MASSCAN_RATE)

	SECRET = viper.GetString(CONFIG_SECRET)
}
