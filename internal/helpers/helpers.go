package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/iamlego/bookingBoost/internal/config"
)

var app *config.AppConfig

func NewHelper(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s \n %s ", err.Error(), debug.Stack())
	app.Errorlog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
