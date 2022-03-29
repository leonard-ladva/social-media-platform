package request

import (
	"forum-test/database"
	"html/template"
	"log"
	"net/http"
	"os"
)

func Delete(w http.ResponseWriter, r *http.Request, user database.User) {
	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	var err error
	var pagedata database.PageData
	pagedata.User = user

	if r.Method == "POST" {

		if err != nil {
			//execute template and return internal server error
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.ExecuteTemplate(w, "err500", pagedata)
			return
		}

		r.ParseForm()
		postid := r.FormValue("postid")
		username := r.FormValue("username")
		image := r.FormValue("postimg")

		if user.Username == username {
			statement, err := database.Db.Prepare("UPDATE post SET postimg = NULL WHERE user_id = ? AND post_id = ? ;")
			if err != nil {
				log.Println(err)
			}
			defer statement.Close()
			statement.Exec(user.ID, postid)

			//also delete the image from server?
			err = os.Remove("public/upload/" + image)
			if err != nil {
				log.Println(err)
			}
		}
		http.Redirect(w, r, "/new-post?id="+postid, http.StatusSeeOther)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	tmpl.ExecuteTemplate(w, "err400", pagedata)
}
