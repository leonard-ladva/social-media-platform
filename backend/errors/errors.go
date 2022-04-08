package errors

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var ErrorLog *log.Logger = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	err = ErrorLog.Output(2, trace)
	if err != nil {
		ServerError(w, err)
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
