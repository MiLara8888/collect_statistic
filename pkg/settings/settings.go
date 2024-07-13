package settings

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Настройка из yaml
type Config struct {
	Host string `env:"HOST,required"`
	Port string `env:"PORT,required"`

	// разрешённые хосты
	HostAlowed string `env:"HOST_ALLOWED,required"`

	Conn Connector
}

// подк. к базам
type DbSettings struct {
	User  string `env:"DB_USER,required"`
	Passw string `env:"DB_PASSW,required"`
	Host  string `env:"DB_HOST,required"`
	Port  uint16 `env:"DB_PORT,required"`
	Sid   string `env:"DB_SID,required"`
}

// сервис
type Connector struct {
	Db DbSettings
}

func GetEnv(key string) (string, error) {
	ret := os.Getenv(key)
	if len(ret) == 0 {
		err := fmt.Errorf("%s env not find ", key)
		return "", err
	}
	return ret, nil
}

func InitEnv() (*Config, error) {
	var err error
	cfg := Config{}
	godotenv.Load()

	err = env.Parse(&cfg) //Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
	return &cfg, nil
}
