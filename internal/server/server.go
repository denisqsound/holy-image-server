package server

import (
	"github.com/denisqsound/holy-image-server/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

func render(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		logrus.Error(err)
	}

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "favicon")
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "img")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "pong")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "robots")
}

func Run(conf config.ConfI) {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)
	if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
		logrus.Fatal(err)
	}
}
