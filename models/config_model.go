package models

type ConfigModel struct {
	StoragePath       string
	AllowedExtensions map[string]bool
}
