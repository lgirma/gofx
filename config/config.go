package config

type AppConfig interface {
	Get(key string, result any) error
	GetEnv() string
}

func GetConfig[T any](conf AppConfig, path string) (T, error) {
	var c T
	err := conf.Get(path, &c)
	return c, err
}

func GetConfigOrDefault[T any](conf AppConfig, path string, defaultVal T) T {
	var c T
	err := conf.Get(path, &c)
	if err != nil {
		return defaultVal
	}
	return c
}
