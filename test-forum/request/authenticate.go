package request

import (
	"fmt"
	"forum-test/database"
	"log"
	"net/http"
	"time"
)

type CustomFunc func(http.ResponseWriter, *http.Request, database.User)

var bucket = make(map[string]int)

func Auth(nextFunction CustomFunc, permission string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//if bucket is 0, it means it's a new ip address. add it to the map. if there are two or more tokens, remove one token. if there is only one token, deny request.
		if bucket[r.RemoteAddr] == 0 {
			bucket[r.RemoteAddr] = 5
		} else if bucket[r.RemoteAddr] >= 2 {
			bucket[r.RemoteAddr]--
		} else if bucket[r.RemoteAddr] == 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "Too many requests, chill down") //replace with a template??
			return
		}

		user := database.GetUserByCookie(w, r)

		if permission == "guest" && user.RoleID != database.GUEST {
			log.Println("Permission denied - Guests only")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if permission == "user" && user.RoleID < database.USER {
			log.Println("Permission denied - Logged in users only")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if permission == "moderator" && user.RoleID < database.MODERATOR {
			log.Println("Permission denied - mods and admins only")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if permission == "admin" && user.RoleID < database.ADMIN {
			log.Println("Permission denied - admins only")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		nextFunction(w, r, user)
	}
}

// add one request token to the bucket every second. maximum on 5 tokens -> user can make up to 4 request burst in one second.
func FillBucket() {
	for {
		for key, v := range bucket {
			if v < 5 {
				bucket[key]++
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// empty token bucket map every 10 minutes
func EmptyBucket() {
	for {
		bucket = make(map[string]int)
		time.Sleep(10 * time.Minute)
	}
}
