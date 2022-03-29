package request

import (
	"forum-test/database"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	var err error
	var pagedata database.PageData
	pagedata.User = user
	pagedata.Categories, err = database.GetCategories()
	if err != nil {
		//execute template and return internal server error
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "err500", pagedata)
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "err404", pagedata)
		return
	}

	tmpl.ExecuteTemplate(w, "index", pagedata)
}
