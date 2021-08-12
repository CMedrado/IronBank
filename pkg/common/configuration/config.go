package configuration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Api      Api
	Database Database
}

type Api struct {
	Port     string `envconfig:"API_PORT" default:"80"`
	LogLevel string `envconfig:"API_LOG_LEVEL" default:"INFO"`
}

type Database struct {
	Protocol     string `envconfig:"DB_PROTOCOL" required:"true"`
	Username     string `envconfig:"DB_USERNAME" required:"true"`
	Secret       string `envconfig:"DB_SECRET" required:"true"`
	Host         string `envconfig:"DB_HOST" required:"true"`
	Port         string `envconfig:"DB_PORT" default:"5432"`
	DatabaseName string `envconfig:"DB_DATABASE" default:"desafio"`
	Options      string `envconfig:"DB_OPTIONS"`
}

func (db Database) GetUrl() string {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		db.Protocol, db.Username, db.Secret, db.Host, db.Port, db.DatabaseName)
	if db.Options != "" {
		url += "?" + db.Options
	}
	return url
}

func LoadConfigs() (Config, error) {
	var config Config
	unprefixed := ""
	err := envconfig.Process(unprefixed, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
