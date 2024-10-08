package configs

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	DBDriver               string `mapstructure:"DB_DRIVER"`
	DBSource               string `mapstructure:"DB_SOURCE"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	FirebaseCredentialPath string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS"`
}

var Module = fx.Options(
	fx.Provide(LoadFirebaseApp),
	fx.Provide(ProvideFirebaseAuth),
	fx.Provide(ProviderFirestore),
)

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
