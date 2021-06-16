package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	BackendHost string 	`mapstructure:"BACKEND_HOST"`
	BackendPort	string	`mapstructure:"BACKEND_PORT"`
	//BackendScheme string `mapstructure:"BACKEND_SCHEME"`

	JwtSignUpEmailSecretKey string `mapstructure:"JWT_SIGN_UP_EMAIL_SECRET_KEY"`
	JwtSignUpEmailAudience string `mapstructure:"JWT_SIGN_UP_EMAIL_AUDIENCE"`

	SmtpServerHost string `mapstructure:"SMTP_SERVER_HOST"`
	SmtpServerPort string `mapstructure:"SMTP_SERVER_PORT"`
	SmtpServerUsername string `mapstructure:"SMTP_SERVER_USERNAME"`
	SmtpServerPassword string `mapstructure:"SMTP_SERVER_PASSWORD"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`

	GinMode string `mapstructure:"GIN_MODE"`

	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresPort string `mapstructure:"POSTGRES_PORT"`
	PostgresDb string `mapstructure:"POSTGRES_DB"`
	PostgresUser string `mapstructure:"POSTGRES_USER"`
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

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// viper will use existing env variable
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}
