type Config struct {
	Port        string `env:"PORT" envDefault:"8080"`
	DatabaseURL string `env:"DATABASE_URL" required:"true"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	RedisURL    string `env:"REDIS_URL"`
}

func Load() *Config {
	cfg := &Config{}
	// Carrega de variáveis de ambiente
	env.Parse(cfg)
	return cfg
}