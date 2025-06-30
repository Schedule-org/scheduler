package domain

type ServiceConfig struct {
	Port                string         `json:"port"`
	RunMigrationEnabled bool           `json:"run_migration_enabled"`
	SwaggerEnabled      bool           `json:"swagger_enabled"`
	Database            DatabaseConfig `json:"database"`
	DevModeEnabled      bool           `json:"dev_mode_enabled"`
	GrafanaEnabled      bool           `json:"grafana_enabled"`
	LoggingEnabled      bool           `json:"logging_enabled"`
}

type DatabaseConfig struct {
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}
