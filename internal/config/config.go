package config


type Config struct {
  Port               string
  DatabaseName       string 
  SessionCookieName  string
}



func loadConfig() (*Config, error) {
	var cfg Config
	// err := envconfig.Process("", &cfg)
	// if err != nil {
	// 	return nil, err
	// }
  cfg.Port = ":8080"
	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
