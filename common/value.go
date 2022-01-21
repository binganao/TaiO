package common

var (
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
)

func InitValue() {

}
