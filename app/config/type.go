package config

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type RedisConfigs struct {
	DefaultRedisConfig *RedisConfig `json:"default"`
	JobRedisConfig     *RedisConfig `json:"job"`
}

type AwsCredentialConfig struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type AwsS3Config struct {
	Bucket string `json:"bucket"`
}

type AwsConfig struct {
	Gateway    string               `json:"gateway"`
	Region     string               `json:"region"`
	Credential *AwsCredentialConfig `json:"credential"`
	S3         *AwsS3Config         `json:"s3"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type AppConfig struct {
	Database *DatabaseConfig `json:"database"`
	Redis    *RedisConfigs   `json:"redis"`
	Server   *ServerConfig   `json:"server"`
	Aws      *AwsConfig      `json:"aws"`
}
