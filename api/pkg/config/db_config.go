package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type DbConfig struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBHostDocker string `mapstructure:"DB_HOST_DOCKER"`
}

var ConfigDB DbConfig

func InitDbConfig(path string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := filepath.Join(currentDir, "pkg/config")
	log.Println(configDir)
	ConfigDB, err = loadDBConfig(path)
	if err != nil {
		log.Fatal(err)
	}
}

func loadDBConfig(path string) (config DbConfig, err error) {
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

func (db *DbConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		db.DBHostDocker, db.DBPort, db.DBUser, db.DBName, db.DBPassword)
}
