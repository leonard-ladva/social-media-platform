package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
)

func MyPosts(w http.ResponseWriter, r *http.Request, user database.User) {

	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	var pagedata database.PageData
	pagedata.User = user
	var err error

	if r.URL.Path != "/myposts" {
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "err404", pagedata)
		return
	}

	pagedata.User = user

	pagedata.Posts, err = database.GetPostsByUserID(user.ID)
	if err != nil {
		log.Println(err)
		//write internal server error header and template
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "err500", pagedata)
		return
	}

	for i := range pagedata.Posts {
		if len(pagedata.Posts[i].Content) > 75 {
			pagedata.Posts[i].Content = pagedata.Posts[i].Content[0:75] + "..."
		}
	}

	tmpl.ExecuteTemplate(w, "category", pagedata)

}
