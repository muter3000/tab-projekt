package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-pg/pg/v10"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/database"
	"github.com/tab-projekt-backend/handlers"
	"github.com/tab-projekt-backend/handlers/administratorzy"
	"github.com/tab-projekt-backend/handlers/kategoria_prawa_jazdy"
	"github.com/tab-projekt-backend/handlers/kierowcy"
	"github.com/tab-projekt-backend/handlers/pracownicy"
	"github.com/tab-projekt-backend/handlers/stanowisko_administracyjne"
)

func main() {
	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatal(err)
	}
	l := hclog.Default()
	l.SetLevel(hclog.Level(int32(logLevel)))
	db, err := database.GetDB()
	if err != nil {
		l.Error("connecting to db", "err", err)
	}
	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	sm := mux.NewRouter()
	subRouters := []handlers.SubRouter{
		pracownicy.NewPracownicy(l, db, "/pracownicy"),
		kierowcy.NewKierowcy(l, db, "/kierowcy"),
		administratorzy.NewAdministratorzy(l, db, "/administracja"),
		stanowisko_administracyjne.NewStanowiskoAdministracyjne(l, db, "/stanowisko_administracyjne"),
		kategoria_prawa_jazdy.NewKategoriaPrawaJazdy(l, db, "/kategoria_prawa_jazdy"),
	}

	for _, sr := range subRouters {
		sr.RegisterSubRouter(sm)
	}
	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	bindAddress := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// create a new server
	s := http.Server{
		Addr:         bindAddress,                                      // configure the bind address
		Handler:      ch(sm),                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  1000 * time.Millisecond,                          // max time to read request from the client
		WriteTimeout: 1000 * time.Millisecond,                          // max time to write response to the client
		IdleTimeout:  1200 * time.Millisecond,                          // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("Starting server", "port", os.Getenv("PORT"))

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = s.Shutdown(ctx)
	if err != nil {
		l.Error("closing server", "err", err)
	}
}
