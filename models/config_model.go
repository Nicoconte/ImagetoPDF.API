package models

type ConfigModel struct {
	StoragePath       string
	RedisUrl          string
	RedisPort         string
	Host              string
	AllowedExtensions map[string]bool
}
