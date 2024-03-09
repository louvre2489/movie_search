package config

import "github.com/caarlos0/env/v10"

type Config struct {
	Env        string `env:"MOVIE_SEARCH_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHOST     string `env:"MOVIE_SEARCH_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"MOVIE_SEARCH_DB_PORT" envDefault:"3306"`
	DBUser     string `env:"MOVIE_SEARCH_DB_USER" envDefault:"todo"`
	DBPassword string `env:"MOVIE_SEARCH_DB_PASSWORD" envDefault:"todo"`
	DBName     string `env:"MOVIE_SEARCH_DB_NAME" envDefault:"todo"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
