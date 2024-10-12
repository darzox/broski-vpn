package config

import (
	"fmt"
	"os"
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
	SupportUserName string `yaml:"supportUserName"`
}

type Service struct {
	config Config
}

func New() (*Service, error) {
	s := &Service{}

	config := Config{
		Token:           getEnv("TOKEN", ""),
		DbUserLogin:     getEnv("DB_USER_LOGIN", "default_user"),
		DbUserPass:      getEnv("DB_USER_PASS", "default_pass"),
		DbHost:          getEnv("DB_HOST", "localhost"),
		DbPort:          getEnvAsInt("DB_PORT", 5432),
		DbName:          getEnv("DB_NAME", "default_db"),
		DbSslMode:       getEnv("DB_SSL_MODE", "disable"),
		VpnUrl:          getEnv("VPN_URL", "http://localhost"),
		MonthPriceInXTR: getEnvAsInt("MONTH_PRICE_IN_XTR", 100),
		SupportUserName: getEnv("SUPPORT_USER_NAME", "broski_support"),
	}
	s.config = config

	return s, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to get an environment variable as an integer
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		var intValue int
		fmt.Sscanf(value, "%d", &intValue)
		return intValue
	}
	return defaultValue
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
