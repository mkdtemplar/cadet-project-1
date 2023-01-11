package configurations

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AppConfig struct {
	TenantID     string `mapstructure:"TENANT_ID"`
	AppId        string `mapstructure:"APP_ID"`
	MSUrl        string `mapstructure:"MS_URL"`
	Crt          string `mapstructure:"CRT"`
	Key          string `mapstructure:"KEY"`
	RootUrl      string `mapstructure:"ROOT_URL"`
	Email        string `mapstructure:"EMAIL"`
	DisplayName  string `mapstructure:"DISPLAY_NAME"`
	ApiPort      string `mapstructure:"API_PORT"`
	UserDelete   string `mapstructure:"USER_DELETE"`
	UserCreate   string `mapstructure:"USER_CREATE"`
	UserPref     string `mapstructure:"USER_PREF"`
	UserPorts    string `mapstructure:"USER_PORTS"`
	ListUserPref string `mapstructure:"LIST_USER_PREF"`
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
