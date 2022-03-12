package config

import (
	"github.com/joho/godotenv"
	"os"
)

type ConfI interface {
	GetPort() string
}

type Conf struct {
	port string
}

func (conf *Conf) GetPort() string {
	return conf.port
}

func New(confPath string) (*Conf, error) {
	err := godotenv.Load(confPath)
	if err != nil {
		return nil, err
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return nil, err
	}
	return &Conf{port: port}, nil
}
