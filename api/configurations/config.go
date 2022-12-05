package configurations

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type AppConfig struct {
	TenantID    string `mapstructure:"TENANT_ID"`
	AppId       string `mapstructure:"APP_ID"`
	MSUrl       string `mapstructure:"MS_URL"`
	Crt         string `mapstructure:"CRT"`
	Key         string `mapstructure:"KEY"`
	RootUrl     string `mapstructure:"ROOT_URL"`
	Email       string `mapstructure:"EMAIL"`
	DisplayName string `mapstructure:"DISPLAY_NAME"`
}

var Config AppConfig

func InitConfig(path string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := filepath.Join(currentDir, "configurations")
	log.Println(configDir)
	Config, err = loadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
