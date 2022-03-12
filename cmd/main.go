package main

import (
	"flag"
	"github.com/denisqsound/holy-image-server/config"
	"github.com/denisqsound/holy-image-server/internal/server"
	"github.com/sirupsen/logrus"
)

var confPath = flag.String("conf-path", "./config/.env", "Path to config .env")

func main() {
	conf, err := config.New(*confPath)
	if err != nil {
		logrus.Fatal(err)
	}
	server.Run(conf)
}
