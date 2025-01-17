package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/iamlego/bookingBoost/internal/config"
	handler "github.com/iamlego/bookingBoost/internal/handlers"
	"github.com/iamlego/bookingBoost/internal/helpers"
	"github.com/iamlego/bookingBoost/internal/models"
	"github.com/iamlego/bookingBoost/internal/render"
)

const port = ":3000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}
	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/About", handler.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", port))
	// _ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() error {
	// put this in session

	gob.Register(models.Reservation{})

	// change this to true in Production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.Errorlog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelper((&app))

	return nil
}
