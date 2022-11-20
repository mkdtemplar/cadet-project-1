package configurations

import "github.com/spf13/viper"

type DbConfig struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPort        string `mapstructure:"DB_PORT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadDBConfig(path string) (config DbConfig, err error) {
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
