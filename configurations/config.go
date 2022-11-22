package configurations

import "github.com/spf13/viper"

type Config struct {
	TenantID string `mapstructure:"TENANT_ID"`
	AppId    string `mapstructure:"APP_ID"`
	MSUrl    string `mapstructure:"MS_URL"`
	Crt      string `mapstructure:"CRT"`
	Key      string `mapstructure:"KEY"`
	RootUrl  string `mapstructure:"ROOT_URL"`
	Email    string `mapstructure:"EMAIL"`
}

func LoadConfig(path string) (config Config, err error) {
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
