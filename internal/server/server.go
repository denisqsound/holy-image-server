package server

import (
	"github.com/denisqsound/holy-image-server/config"
	"github.com/denisqsound/holy-image-server/pkg/img"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func render(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		logrus.Error(err)
	}

}

func logoHandler(w http.ResponseWriter, r *http.Request) {
	buffer, err := img.GetLogo()
	if err != nil {
		logrus.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Conent-Type", "image/octet-stream")
	w.Write(buffer)
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
	http.HandleFunc("/logo", logoHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots.txt", robotsHandler)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":"+conf.GetPort(), nil); err != nil {
			logrus.Fatal(err)
		}
	}()

	signalValue := <-sigs
	signal.Stop(sigs)
	logrus.Println("stop signal: ", signalValue)
}
