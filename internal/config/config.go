package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Grpc struct {
		Port               string `mapstructure:"port"`
		AuthServiceAddress string `mapstructure:"auth_service_address"`
	} `mapstructure:"grpc"`

	Logging struct {
		IsProduction bool   `mapstructure:"isProduction"`
		VectorURL    string `mapstructure:"vectorURL"`
	} `mapstructure:"logging"`

	Postgres struct {
		User     string `mapstructure:"user"`
		Name     string `mapstructure:"name"`
		Password string `mapstructure:"password"`
		Address  string `mapstructure:"address"`
	} `mapstructure:"postgres"`

	AIModel struct {
		ModelUri    string  `mapstructure:"model_uri"`
		MaxTokens   int     `mapstructure:"max_tokens"`
		Temperature float32 `mapstructure:"temperature"`
		AuthToken   string  `mapstructure:"auth_token"`
		Address     string  `mapstructure:"address"`
	} `mapstructure:"ai_model"`
}

var Cfg Config

func LoadConfig(path string) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("Config not found, used env and default values")
		} else {
			log.Error().Msg("Failed read config file")
		}
	}

	for _, k := range v.AllKeys() {
		value := v.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envVar := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
			envValue := os.Getenv(envVar)
			if envValue != "" {
				v.Set(k, envValue)
			}
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Error().Msg("Failed unmarshal config")
	}

	if err := validateConfig(&cfg); err != nil {
		log.Fatal().Msg("Not valid config")
	}

	Cfg = cfg
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("grpc.port", "8000")
	v.SetDefault("grpc.auth_service_address", "")

	v.SetDefault("postgres.address", "http://localhost:5432")
	v.SetDefault("postgres.user", "admin")
	v.SetDefault("postgres.password", "password")
	v.SetDefault("postgres.name", "db")

	v.SetDefault("logging.isProduction", false)
	v.SetDefault("logging.vectorURL", "http://vector:9880")

	v.SetDefault("ai_model.model_uri", "")
	v.SetDefault("ai_model.max_tokens", "0")
	v.SetDefault("ai_model.temperature", "0.0")
	v.SetDefault("ai_model.auth_token", "")
	v.SetDefault("ai_model.address", "")
}

func validateConfig(cfg *Config) error {
	return nil
}
