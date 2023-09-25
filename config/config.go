package config

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type PostgresqlConfigYaml struct {
	Host            string        `yaml:"host"`
	Port            uint          `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Database        string        `yaml:"database"`
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type ServerConfigYaml struct {
	EnableHTTPS  bool                 `yaml:"enable_https"`
	Port         uint                 `yaml:"port"`
	MigrationDir string               `yaml:"migration_dir"`
	Postgresql   PostgresqlConfigYaml `yaml:"postgresql"`
}

var ServerConfig = &ServerConfigYaml{}

func CreateServerConfig() error {

	pgHost, ok := os.LookupEnv(EnvPgHost)
	if ok {
		ServerConfig.Postgresql.Host = pgHost
	}
	pgPort, ok := os.LookupEnv(EnvPgPort)
	if ok {
		pgPort_, err := strconv.Atoi(pgPort)
		if err != nil {
			return errors.Wrap(err, "convert port from env to int")
		}
		ServerConfig.Postgresql.Port = uint(pgPort_)
	} else {
		ServerConfig.Postgresql.Port = 5432
	}
	pgUser, ok := os.LookupEnv(EnvPgUser)
	if ok {
		ServerConfig.Postgresql.User = pgUser
	}
	pgPassword, ok := os.LookupEnv(EnvPgPassword)
	if ok {
		ServerConfig.Postgresql.Password = pgPassword
	}
	pgDatabase, ok := os.LookupEnv(EnvPgDatabase)
	if ok {
		ServerConfig.Postgresql.Database = pgDatabase
	}
	pgSSLMode, ok := os.LookupEnv(EnvPgSSLMode)
	if ok {
		ServerConfig.Postgresql.SSLMode = pgSSLMode
	}
	ServerConfig.Postgresql.MaxOpenConns = 20
	pgMaxOpenConns, ok := os.LookupEnv(EnvPgMaxOpenConns)
	if ok {
		pgMaxOpenConns_, err := strconv.Atoi(pgMaxOpenConns)
		if err != nil {
			return errors.Wrap(err, "convert PG_MAX_OPEN_CONNS from env to int")
		}
		ServerConfig.Postgresql.MaxOpenConns = pgMaxOpenConns_
	}
	ServerConfig.Postgresql.MaxIdleConns = 20
	pgMaxIdleConns, ok := os.LookupEnv(EnvPgMaxIdleConns)
	if ok {
		pgMaxIdleConns_, err := strconv.Atoi(pgMaxIdleConns)
		if err != nil {
			return errors.Wrap(err, "convert PG_MAX_IDLE_CONNS from env to int")
		}
		ServerConfig.Postgresql.MaxIdleConns = pgMaxIdleConns_
	}
	ServerConfig.Postgresql.ConnMaxLifetime = 15 * time.Minute
	pgConnMaxLifetime, ok := os.LookupEnv(EnvPgConnsMaxLifetime)
	if ok {
		pgConnMaxLifetime_, err := time.ParseDuration(pgConnMaxLifetime)
		if err != nil {
			return errors.Wrap(err, "convert PG_CONN_MAX_LIFETIME from env to time.Duration")
		}
		ServerConfig.Postgresql.ConnMaxLifetime = pgConnMaxLifetime_
	}
	migrationDir, ok := os.LookupEnv(EnvMigrationDir)
	if ok {
		ServerConfig.MigrationDir = migrationDir
	}

	if ServerConfig.Port == 0 {
		ServerConfig.Port = 5000
	}

	return nil
}
