package config

type Environment = string

type ConfigSourceType = string

const (
	EnvDefault    Environment = ""
	EnvDev        Environment = "dev"
	EnvReview     Environment = "review"
	EnvProduction Environment = "production"
)

const (
	ConfigSourceYaml    ConfigSourceType = "yaml"
	ConfigSourceEnvVars ConfigSourceType = "env"
	ConfigSourceJson    ConfigSourceType = "json"
	ConfigSourceToml    ConfigSourceType = "toml"
)
