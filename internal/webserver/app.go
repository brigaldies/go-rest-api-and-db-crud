package webserver

import (
	"database/sql"
	"fmt"
	"github.com/brigaldies/go-rest-api-and-db-crud/internal/api"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // Postgres database SQL driver
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	fmt.Println("Initializing...")
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	api.InitializeRoutes(a.DB, a.Router)
}

func (a *App) Run(addr string) {
	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}
