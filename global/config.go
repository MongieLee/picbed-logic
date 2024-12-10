package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type MySqlConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

var ProConfig struct {
	Mysql MySqlConfig `mapstructure:"mysql"`
}

func InitViperConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.Unmarshal(&ProConfig)
		if err != nil {
			log.Fatalf("Viper反解析Json失败，错误信息：%v", err)
		}
	})
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Viper读取配置失败，%v.\n", err)
	}
	err = viper.Unmarshal(&ProConfig)
	if err != nil {
		log.Fatalf("Viper反解析Json失败，错误信息：%v", err)
	}
}
