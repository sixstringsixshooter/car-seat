package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sixstringsixshooter/car-seat/config"
	"go.uber.org/zap"
)

const (
	baseV1          = "/api/v1"
	checkStatusPath = baseV1 + "/status"
)

type controller struct {
	log    *zap.Logger
	router *mux.Router
}

func newController(conf *config.Config) (*controller, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	c := controller{
		log:    l,
		router: mux.NewRouter(),
	}

	if on, ok := conf.Versions["v1"]; ok && on {
		l.Info("adding v1 routes...")
		c.addV1Routes(c.router.PathPrefix(baseV1).Subrouter())
	}
	return &c, nil
}

func (c *controller) addV1Routes(router *mux.Router) {
	c.router.HandleFunc(checkStatusPath, c.checkStatusHandler()).Methods(http.MethodGet)
}

func (c *controller) checkStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	}
}
