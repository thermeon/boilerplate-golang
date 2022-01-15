package configs

import (
	"context"
	"github.com/thermeon/boilerplate-golang/log"
	"github.com/thermeon/boilerplate-golang/utils"
)

type Config struct {
	AdapterConfigPath            string `env:"ADAPTER_CONFIG_PATH"`
	AdapterSecretPath            string `env:"ADAPTER_SECRET_PATH"`
	ApplicationName              string `env:"APPLICATION_NAME"`
	DebugCachePath               string `env:"DEBUG_CACHE_PATH"`
	DebugEnable                  bool   `env:"DEBUG_ENABLE;optional"`
	DebugMemoryStore             bool   `env:"DEBUG_MEMORYSTORE;optional"`
	DebugSuppressMetrics         bool   `env:"DEBUG_SUPPRESS_METRICS;optional"`
	EventFilters                 string `env:"EVENT_FILTERS"`
	GoogleApplicationCredentials string `env:"GOOGLE_APPLICATION_CREDENTIALS"`
	I18NConfigPath               string `env:"I18N_CONFIG_PATH"`
	PostgresClientCert           string `env:"POSTGRES_CLIENTCERT"`
	PostgresClientKey            string `env:"POSTGRES_CLIENTKEY"`
	PostgresConnType             string `env:"POSTGRES_CONNTYPE"`
	PostgresDbname               string `env:"POSTGRES_DBNAME"`
	PostgresHost                 string `env:"POSTGRES_HOST"`
	PostgresPassword             string `env:"POSTGRES_PASSWORD"`
	PostgresPort                 string `env:"POSTGRES_PORT"`
	PostgresServerCA             string `env:"POSTGRES_SERVERCA"`
	PostgresUser                 string `env:"POSTGRES_USER"`
	ProjectName                  string `env:"PROJECT_NAME"`
	ReceiptTemplatePath          string `env:"RECEIPT_TEMPLATE_PATH"`
	RSAPublicKeyPath             string `env:"RSA_PUBLIC_KEY_PATH"`
	SecurityAudience             string `env:"SECURITY_AUDIENCE"`
	TimezoneMapPath              string `env:"TIMEZONE_MAP_PATH"`
}

// Init read the environment variables and relevant yaml files.
// Validate the configurations at the end.
func (c *Config) Init() {
	ctx := context.Background()
	if err := utils.NewEnvConfigReader().Read(c); err != nil {
		log.Fatal(ctx, "Unable to read environment variables. Error: %s", err)
	}
	c.validate()
}

// validate validations for custom configurations.
func (c *Config) validate() {

}
