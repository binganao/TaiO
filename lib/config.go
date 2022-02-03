package lib

import (
	"github.com/binganao/TaiO/common"
	"github.com/binganao/TaiO/pkg/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// node
	viper.SetDefault(common.CONFIG_NODE_TYPE, "")
	viper.SetDefault(common.CONFIG_DATA_ADDR, "http://127.0.0.1:8080")

	// database
	viper.SetDefault(common.CONFIG_DATABASE_DRIVER, "mysql")
	viper.SetDefault(common.CONFIG_DATABASE_HOST, "localhost")
	viper.SetDefault(common.CONFIG_DATABASE_PORT, "3306")
	viper.SetDefault(common.CONFIG_DATABASE_USERNAME, "root")
	viper.SetDefault(common.CONFIG_DATABASE_PASSWORD, "root")
	viper.SetDefault(common.CONFIG_DATABASE_DATABASE, "TaiO")
	viper.SetDefault(common.CONFIG_DATABASE_CHARSET, "utf8mb4")

	// masscan
	viper.SetDefault(common.CONFIG_MASSCAN_RATE, 100)

	// add secret
	viper.SetDefault(common.CONFIG_SECRET, "")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("读取配置文件失败，若为第一次使用，请重启 TaiO")
	}

	err = viper.WriteConfigAs("config.toml")
	if err != nil {
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("配置文件已更新！")
		common.InitValue()
	})
}
