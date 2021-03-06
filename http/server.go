package http

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	router := httprouter.New()

	router.GET("/", versionHandler)
	router.GET("/version", versionHandler)

	router.GET("/games/:game/gif", handleGIFGame)

	router.GET("/games/:game/frames/:frame/ascii", handleASCIIFrame)
	router.GET("/games/:game/frames/:frame/gif", handleGIFFrame)

	router.GET("/healthz/alive", handleAlive)
	router.GET("/healthz/ready", handleReady)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = ":8000"
	}
	log.WithField("port", port).Info("http server listening")
	if err := http.ListenAndServe(port, router); err != nil {
		log.WithError(err).WithField("port", port).Error("error while trying to listen on port")
	}
}
