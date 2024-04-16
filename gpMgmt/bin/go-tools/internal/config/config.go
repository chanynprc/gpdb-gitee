package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	SetName(string)
	Load(string) error

	GetDatabaseConfig() DatabaseConfig
	GetArtifactConfig() ArtifactConfig
	GetInfraConfig() InfraConfig
}

func New() Config {
	return &appConfig{
		configName: "gp.conf",
	}
}

type appConfig struct {
	configName string
	// TODO : embed hub Config

	Database *Database `json:"database"`
	Artifact *Artifact `json:"artifact"`
	Infra    *Infra    `json:"infra"`
}

func (conf *appConfig) SetName(configName string) {
	conf.configName = configName
}

func (conf *appConfig) Load(configPath string) error {
	parser := viper.New()
	parser.SetConfigName(conf.configName)
	parser.SetConfigType("json")

	parser.AddConfigPath(configPath)

	conf.setDefaults(parser)
	err := parser.ReadInConfig()
	if err != nil {
		return err
	}

	err = parser.Unmarshal(conf)
	if err != nil {
		return err
	}
	return nil
}

func (conf *appConfig) GetDatabaseConfig() DatabaseConfig {
	return conf.Database
}

func (conf *appConfig) GetArtifactConfig() ArtifactConfig {
	return conf.Artifact
}

func (conf *appConfig) GetInfraConfig() InfraConfig {
	return conf.Infra
}

func (conf *appConfig) setDefaults(parser *viper.Viper) *viper.Viper {
	parser.SetDefault("Infra.RequestPort", 4506)
	parser.SetDefault("Infra.PublishPort", 4505)
	parser.SetDefault("Infra.Coordinator.HostName", "cdw")
	parser.SetDefault("Database.Admin.Name", "gpadmin")

	return parser
}
