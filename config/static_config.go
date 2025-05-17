package config

type StaticConfig struct {
	configMap map[string]any
}

func (s *StaticConfig) Get(key string, result any) error {
	result = s.configMap[key]
	return nil
}

func (s *StaticConfig) GetEnv() string {
	return EnvDev
}

func NewStaticConfig(src map[string]any) AppConfig {
	return &StaticConfig{
		configMap: src,
	}
}
