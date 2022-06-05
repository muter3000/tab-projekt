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

	"github.com/go-openapi/runtime/middleware"

	"github.com/tab-projekt-backend/database/psql"
	"github.com/tab-projekt-backend/handlers"
	"github.com/tab-projekt-backend/handlers/server/administratorzy"
	"github.com/tab-projekt-backend/handlers/server/bledy"
	"github.com/tab-projekt-backend/handlers/server/kategoria_prawa_jazdy"
	"github.com/tab-projekt-backend/handlers/server/kierowcy"
	"github.com/tab-projekt-backend/handlers/server/kursy"
	"github.com/tab-projekt-backend/handlers/server/marki"
	"github.com/tab-projekt-backend/handlers/server/pojazdy"
	"github.com/tab-projekt-backend/handlers/server/pojazdy_ciezarowe"
	"github.com/tab-projekt-backend/handlers/server/pracownicy"
	"github.com/tab-projekt-backend/handlers/server/stanowisko_administracyjne"
	"github.com/tab-projekt-backend/handlers/server/trasy"
	"github.com/tab-projekt-backend/schemas"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/rs/cors"
)

// type dbLogger struct{}

// func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
// 	return c, nil
// }

// func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
// 	value, _ := q.FormattedQuery()
// 	fmt.Println(string(value))
// 	return nil
// }

func main() {
	logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatal(err)
	}
	l := hclog.Default()
	l.SetLevel(hclog.Level(int32(logLevel)))
	db, err := psql.GetDB()
	if err != nil {
		l.Error("connecting to db", "err", err)
	}
	defer func(db *pg.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// db.AddQueryHook(dbLogger{})

	orm.RegisterTable((*schemas.KategoriaKierowcy)(nil))

	sm := mux.NewRouter()

	g := sm.Methods(http.MethodGet).Subrouter()
	sh := middleware.Redoc(middleware.RedocOpts{SpecURL: "/swagger.yaml"}, nil)
	g.Handle("/docs", sh)
	g.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	subRouters := []handlers.SubRouter{
		pracownicy.NewPracownicy(l, db, "/pracownicy"),
		kierowcy.NewKierowcy(l, db, "/kierowcy"),
		administratorzy.NewAdministratorzy(l, db, "/administracja"),
		stanowisko_administracyjne.NewStanowiskoAdministracyjne(l, db, "/stanowiska_administracyjne"),
		kategoria_prawa_jazdy.NewKategoriaPrawaJazdy(l, db, "/kategoria_prawa_jazdy"),
		trasy.NewTrasy(l, db, "/trasy"),
		pojazdy.NewPojazdy(l, db, "/pojazdy"),
		marki.NewMarki(l, db, "/marki"),
		pojazdy_ciezarowe.NewPojazdyCiezarowe(l, db, "/pojazdy_ciezarowe"),
		kursy.NewKursy(l, db, "/kursy"),
		bledy.NewBledy(l, db, "/bledy"),
	}

	for _, sr := range subRouters {
		sr.RegisterSubRouter(sm)
	}

	handler := cors.AllowAll().Handler(sm)

	bindAddress := fmt.Sprintf(":%s", os.Getenv("PORT"))

	// create a new server
	s := http.Server{
		Addr:         bindAddress,                                      // configure the bind address
		Handler:      handler,                                          // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  500 * time.Millisecond,                           // max time to read request from the client
		WriteTimeout: 1000 * time.Millisecond,                          // max time to write response to the client
		IdleTimeout:  12000 * time.Millisecond,                         // max time for connections using TCP Keep-Alive
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
