package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AppConfig struct {
	TenantID      string `mapstructure:"TENANT_ID"`
	AppId         string `mapstructure:"APP_ID"`
	MSUrl         string `mapstructure:"MS_URL"`
	Crt           string `mapstructure:"CRT"`
	Key           string `mapstructure:"KEY"`
	RootUrl       string `mapstructure:"ROOT_URL"`
	Email         string `mapstructure:"EMAIL"`
	DisplayName   string `mapstructure:"DISPLAY_NAME"`
	ApiPort       string `mapstructure:"API_PORT"`
	UserDelete    string `mapstructure:"USER_DELETE"`
	UserCreate    string `mapstructure:"USER_CREATE"`
	UserPref      string `mapstructure:"USER_PREF"`
	UserPorts     string `mapstructure:"USER_PORTS"`
	ListUserPref  string `mapstructure:"LIST_USER_PREF"`
	UserId        string `mapstructure:"USER_GET"`
	UserPrefPorts string `mapstructure:"USER_PREF_PORTS"`
	MapsKey       string `mapstructure:"GOOGLE_API_KEY"`
	PortName      string `mapstructure:"PORT_NAME"`
}

var Config AppConfig

func InitConfig(path string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := filepath.Join(currentDir, "pkg/config")
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
