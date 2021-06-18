package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yerassyldanay/invest/utils/helper"
)

type Config struct {
	LoadEnvFile bool `mapstructure:"LOAD_ENV_FILE"`

	BackendHost string `mapstructure:"BACKEND_HOST"`
	BackendPort string `mapstructure:"BACKEND_PORT"`
	//BackendScheme string `mapstructure:"BACKEND_SCHEME"`

	JwtSignUpEmailSecretKey string `mapstructure:"JWT_SIGN_UP_EMAIL_SECRET_KEY"`
	JwtSignUpEmailAudience  string `mapstructure:"JWT_SIGN_UP_EMAIL_AUDIENCE"`

	SmtpServerHost     string `mapstructure:"SMTP_SERVER_HOST"`
	SmtpServerPort     string `mapstructure:"SMTP_SERVER_PORT"`
	SmtpServerUsername string `mapstructure:"SMTP_SERVER_USERNAME"`
	SmtpServerPassword string `mapstructure:"SMTP_SERVER_PASSWORD"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`

	GinMode string `mapstructure:"GIN_MODE"`

	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`

	//SMSApiKey string `mapstructure:"SMS_API_KEY"`
}

func (c *Config) GetDatabaseSource() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresDb,
		c.PostgresPassword)
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("local")

	viper.SetDefault("LOAD_ENV_FILE", false)

	viper.SetDefault("BACKEND_HOST", "0.0.0.0")
	viper.SetDefault("BACKEND_PORT", "7000")

	viper.SetDefault("JWT_SIGN_UP_EMAIL_SECRET_KEY", "")
	viper.SetDefault("JWT_SIGN_UP_EMAIL_AUDIENCE", "")

	viper.SetDefault("SMTP_SERVER_HOST", "")
	viper.SetDefault("SMTP_SERVER_PORT", "")
	viper.SetDefault("SMTP_SERVER_USERNAME", "")
	viper.SetDefault("SMTP_SERVER_PASSWORD", "")

	viper.SetDefault("REDIS_HOST", "invest_redis")
	viper.SetDefault("REDIS_PORT", "6379")

	viper.SetDefault("GIN_MODE", "")

	viper.SetDefault("POSTGRES_HOST", "invest_postgres")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_DB", "simple")
	viper.SetDefault("POSTGRES_USER", "simple")
	viper.SetDefault("POSTGRES_PASSWORD", "simple")

	err = viper.Unmarshal(&config)
	helper.IfErrorPanic(err)

	if config.LoadEnvFile {
		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		// viper will use existing env variable
		viper.AutomaticEnv()

		err = viper.Unmarshal(&config)
	}
	return
}
