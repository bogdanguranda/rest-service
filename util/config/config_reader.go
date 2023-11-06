package config

type ConfigReader interface {
	ReadConfig(filename string) (Config, error)
}
