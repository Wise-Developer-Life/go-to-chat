package config

func getDefaultConfig() *AppConfig {
	return &AppConfig{
		Database: &DatabaseConfig{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "",
			DBName:   "go_to_chat",
		},
		Redis: &RedisConfigs{
			DefaultRedisConfig: &RedisConfig{
				Host:     "localhost",
				Port:     6378,
				Password: "",
				DB:       0,
			},
			JobRedisConfig: &RedisConfig{
				Host:     "localhost",
				Port:     6378,
				Password: "",
				DB:       1,
			},
		},
		Server: &ServerConfig{
			Port: 8080,
		},
	}
}
