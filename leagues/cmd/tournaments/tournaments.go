package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amauryprieto/pkg/tournaments"
	"github.com/go-kit/kit/log"
)

type config struct {
	httpAddr string
	dsconfig string
	logger   log.Logger
}

func setup() *config {
	l := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	c := config{
		logger: log.With(l, "ts", log.DefaultTimestampUTC),
	}

	flag.StringVar(&c.httpAddr, "addr", "localhost:8080", "HTTP address")
	flag.StringVar(&c.dsconfig, "dsconfig", "host=localhost port=5432 user=pdm_dev dbname=pdm_dev password=pdm_dev sslmode=disable", "Datasource config")

	return &c
}

func main() {
	c := setup()
	db, err := sql.Open("postgres", c.dsconfig)
	if err != nil {
		c.logger.Log(err)
	}

	r := tournaments.NewUserRepository(db, c.logger)
	s := tournaments.NewService(r, c.logger)

	mux := http.NewServeMux()
	mux.Handle("/tournaments/", tournaments.MakeHandler(s, c.logger))

	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)
	go func() {
		c.logger.Log("transport", "http", "address", c.httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(c.httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	c.logger.Log("terminated", <-errs)

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
