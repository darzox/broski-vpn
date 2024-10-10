package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const configFile = "data/config.yaml"

type Config struct {
	Token           string `yaml:"token"`
	DbUserLogin     string `yaml:"postgresUserLogin"`
	DbUserPass      string `yaml:"postgresUserPass"`
	DbHost          string `yaml:"postgresHost"`
	DbPort          int    `yaml:"postgresPort"`
	DbName          string `yaml:"postgresDBName"`
	DbSslMode       string `yaml:"postgresSslMode"`
	VpnUrl          string `yaml:"vpnUrl"`
	MonthPriceInXTR int    `yaml:"monthPriceInXTR"`
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	rawYaml, err := os.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYaml, &s.config)
	if err != nil {
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return s, nil
}

func (s *Service) Token() string {
	return s.config.Token
}

func (s *Service) VpnUrl() string {
	return s.config.VpnUrl
}

func (s *Service) PostgresDBName() string {
	return s.config.DbName
}

func (s *Service) PostgresDBUserLogin() string {
	return s.config.DbUserLogin
}

func (s *Service) PostgresUserPass() string {
	return s.config.DbUserPass
}

func (s *Service) PostgresHost() string {
	return s.config.DbHost
}

func (s *Service) PostgresPort() int {
	return s.config.DbPort
}

func (s *Service) PostgresSslMode() string {
	return s.config.DbSslMode
}

func (s *Service) MonthPriceInXTR() int {
	return s.config.MonthPriceInXTR
}
