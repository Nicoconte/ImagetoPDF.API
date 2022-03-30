package models

type ConfigModel struct {
	StoragePath       string `env:"STORAGE_PATH"`
	RedisUrl          string `env:"REDIS_URL"`
	Port              string `env:"PORT"`
	AllowedExtensions string `env:"ALLOWED_EXTENSIONS"`
}
