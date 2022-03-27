package models

type ConfigModel struct {
	StoragePath       string
	CommandExecutor   string
	RedisConnection   string
	Host              string
	AllowedExtensions map[string]bool
}
