package config

const (
	EnvIsSaaS           = "IS_SAAS"
	EnvSaasDomainSuffix = "SAAS_DOMAIN_SUFFIX"

	EnvPgHost             = "PG_HOST"
	EnvPgPort             = "PG_PORT"
	EnvPgUser             = "PG_USER"
	EnvPgPassword         = "PG_PASSWORD"
	EnvPgDatabase         = "PG_DATABASE"
	EnvPgSSLMode          = "PG_SSLMODE"
	EnvPgMaxOpenConns     = "PG_MAX_OPEN_CONNS"
	EnvPgMaxIdleConns     = "PG_MAX_IDLE_CONNS"
	EnvPgConnsMaxLifetime = "PG_CONN_MAX_LIFETIME"

	EnvMigrationDir = "MIGRATION_DIR"

	EnvServerPort = "SERVER_PORT"

	EnvDockerImageBuilderPrivileged = "DOCKER_IMAGE_BUILDER_PRIVILEGED"

	EnvReadHeaderTimeout = "READ_HEADER_TIMEOUT"

	EnvTransmissionStrategy = "TRANSMISSION_STRATEGY"
)
