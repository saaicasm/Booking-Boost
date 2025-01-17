package handler

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/iamlego/bookingBoost/internal/config"
	"github.com/iamlego/bookingBoost/internal/models"
	"github.com/iamlego/bookingBoost/internal/render"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})

	// change this to true in Production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.Errorlog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(WriteToDConsole)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/About", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)
	mux.Get("/search-availability", Repo.Availability)
	mux.Get("/make-reservation", Repo.Reservation)
	mux.Get("/contact", Repo.Contact)

	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Post("/search-availability-json", Repo.PostAvailabilityJSON)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static"))
	// log.Println("The content of file server are : ", fileServer) // it is &{./static}

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func WriteToDConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
