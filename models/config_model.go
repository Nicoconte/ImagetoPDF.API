package models

type ConfigModel struct {
	Env               string `env:"APP_ENV"`
	StoragePath       string `env:"STORAGE_PATH"`
	RedisUrl          string `env:"REDIS_URL"`
	Port              string `env:"PORT"`
	AllowedExtensions string `env:"ALLOWED_EXTENSIONS"`
	Host              string `env:"HOST"`
}
