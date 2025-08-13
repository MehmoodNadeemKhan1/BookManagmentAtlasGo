package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"POSTGRESS_HOST"`
	DBUsername string `mapstructure:"POSTGRESS_USER"`
	DBPassword string `mapstructure:"POSTGRESS_PASSWORD"`
	DBName     string `mapstructure:"POSTGRESS_DB"`
	DBPort     string `mapstructure:"POSTGRESS_PORT"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	viper.Unmarshal(&config)
	return

}
