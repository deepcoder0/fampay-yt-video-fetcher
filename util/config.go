package util

import "github.com/spf13/viper"

type Config struct {
	ServerAddress      string      `mapstructure:"SERVER_ADDRESS"`
	MongoURI           string      `mapstructure:"MONGODB_URI"` 
	APIKeysYT          []string    `mapstructure:"API_KEYS_YT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName(".env")
    viper.SetConfigType("env")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}