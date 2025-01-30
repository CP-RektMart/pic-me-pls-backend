package postgres

import "fmt"

type Config struct {
	Host     string `env:"HOST"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DBNAME"`
	Port     int    `env:"PORT"`
	SSLMode  string `env:"SSLMODE"`
}

func (c Config) String() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
	return dsn
}
