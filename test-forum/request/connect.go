package request

import (
	"encoding/json"
	"fmt"
	"forum-test/database"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const GOOGLETOKENURL = "https://oauth2.googleapis.com/token"
const GHTOKENURL = "https://github.com/login/oauth/access_token"
const FBTOKENURL = "https://graph.facebook.com/v12.0/oauth/access_token?"

const CALLBACK_URI = "https://localhost:8000/connect"

var (
	GCLIENT_ID,
	GOOGLE_SECRET,
	FBCLIENT_ID,
	FB_SECRET,
	GHCLIENT_ID,
	GH_SECRET string
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	IDToken     string `json:"id_token"`
}

type UserData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type FBID struct {
	Data struct {
		UserID string `json:"user_id"`
	} `json:"data"`
}

var GitUserData struct {
	Login string `json:"login"`
}

var GitUserEmails []struct {
	Email string `json:"email"`
}

var (
	data       url.Values
	code       []string
	IDTokenUrl string
	userdata   UserData
	res        TokenResponse
)

func ConnectOauth(w http.ResponseWriter, r *http.Request, user database.User) {
	var pagedata database.PageData
	pagedata.User = user
	//tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))

	// Get client ID-s and secrets from OS environment variables
	GCLIENT_ID = os.Getenv("GCLIENT_ID")       //Google client id
	FBCLIENT_ID = os.Getenv("FBCLIENT_ID")     //Facebook app id
	GHCLIENT_ID = os.Getenv("GHCLIENT_ID")     //Github client/app id
	FB_SECRET = os.Getenv("FB_SECRET")         //Facebook secret (DONT SHARE!)
	GOOGLE_SECRET = os.Getenv("GOOGLE_SECRET") //Google secret
	GH_SECRET = os.Getenv("GH_SECRET")         //Github secret

	log.Println(GCLIENT_ID)

	//Data for HTML template (OAuth links)
	pagedata.Authdata.GoogleClientID = GCLIENT_ID
	pagedata.Authdata.FacebookClientID = FBCLIENT_ID
	pagedata.Authdata.GitHubClientID = GHCLIENT_ID
	pagedata.CALLBACK_URI = CALLBACK_URI

	//get the OAuth provider from url query
	provider := r.URL.Query()["state"][0]

	//get the code from oauth request
	code = r.URL.Query()["code"]

	if provider == "google" {

		IDTokenUrl = "https://oauth2.googleapis.com/tokeninfo?id_token="
		data = url.Values{
			"code":          {code[0]},
			"client_id":     {GCLIENT_ID},
			"client_secret": {GOOGLE_SECRET},
			"redirect_uri":  {CALLBACK_URI},
			"grant_type":    {"authorization_code"},
		}
		log.Println("OAuth Provider:", provider)
		log.Println("OAuth code:", code[0])

		//exchange code for access token
		resp, err := http.PostForm(GOOGLETOKENURL, data)
		check(err, w, r, user)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		check(err, w, r, user)

		err = json.Unmarshal(body, &res)
		check(err, w, r, user)

		//get the contents of IDToken
		resp, err = http.Get(IDTokenUrl + res.IDToken)
		check(err, w, r, user)

		body, err = ioutil.ReadAll(resp.Body)
		check(err, w, r, user)
		defer resp.Body.Close()
		//read the response into userdata struct
		err = json.Unmarshal(body, &userdata)
		check(err, w, r, user)
		log.Println(userdata)

	}

	if provider == "facebook" {
		//get the code from oauth
		IDTokenUrl = "https://graph.facebook.com/debug_token?"
		data = url.Values{
			"code":          {code[0]},
			"client_id":     {FBCLIENT_ID},
			"client_secret": {FB_SECRET},
			"redirect_uri":  {CALLBACK_URI},
			"grant_type":    {"authorization_code"},
		}

		log.Println("OAuth Provider:", provider)
		log.Println("OAuth code:", code[0])

		//get access token from FB
		resp, err := http.PostForm(FBTOKENURL, data)
		check(err, w, r, user)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		check(err, w, r, user)

		//read the access token and save it to a struct
		err = json.Unmarshal(body, &res)
		check(err, w, r, user)

		//inspect the token and get userID
		resp, err = http.Get(IDTokenUrl + "input_token=" + res.AccessToken + "&access_token=" + FBCLIENT_ID + "|" + FB_SECRET)
		check(err, w, r, user)

		body, err = ioutil.ReadAll(resp.Body)
		check(err, w, r, user)
		defer resp.Body.Close()

		var res2 FBID
		err = json.Unmarshal(body, &res2)
		check(err, w, r, user)

		//get user e-mail and name by userID and access token

		resp, err = http.Get("https://graph.facebook.com/v12.0/" + res2.Data.UserID + "/?fields=email%2Cname%2Cpicture&access_token=" + res.AccessToken)
		check(err, w, r, user)

		body, err = ioutil.ReadAll(resp.Body)
		check(err, w, r, user)
		defer resp.Body.Close()

		err = json.Unmarshal(body, &userdata)
		check(err, w, r, user)
	}

	if provider == "github" {

		IDTokenUrl = "https://api.github.com/user"
		data = url.Values{
			"code":          {code[0]},
			"client_id":     {GHCLIENT_ID},
			"client_secret": {GH_SECRET},
			"redirect_uri":  {CALLBACK_URI},
			"grant_type":    {"authorization_code"},
		}
		log.Println("OAuth Provider:", provider)
		log.Println("OAuth code:", code[0])

		log.Println("Github secret:", GH_SECRET)

		//exchange code for access token
		resp, err := http.PostForm(GHTOKENURL, data)
		check(err, w, r, user)
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		check(err, w, r, user)

		res.AccessToken = strings.Split(string(body), "=")[1]
		res.AccessToken = strings.Split(res.AccessToken, "&")[0]
		if len(res.AccessToken) < 1 {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal server error!")
		}
		log.Println(string(body))
		log.Println("Access token:", res.AccessToken)

		if res.AccessToken == "bad_verification_code" {
			log.Println("Github: Bad access token!")
			//internal server error and return
			fmt.Fprintf(w, "Something went wrong with authentication, sorry!")
		}

		//get the user data
		client := &http.Client{}

		req, err := http.NewRequest(http.MethodGet, IDTokenUrl, nil)
		check(err, w, r, user)

		req.Header.Set("Authorization", "token "+res.AccessToken)
		req.Header.Set("Accept", "application/vnd.github.v3+json")
		log.Println(req)
		resp, err = client.Do(req)

		check(err, w, r, user)

		body, err = ioutil.ReadAll(resp.Body)
		check(err, w, r, user)
		defer resp.Body.Close()

		log.Println(string(body))
		//read the response into userdata struct
		err = json.Unmarshal(body, &GitUserData)
		log.Println(GitUserData)
		check(err, w, r, user)

		req, err = http.NewRequest(http.MethodGet, IDTokenUrl+"/emails", nil)
		check(err, w, r, user)

		req.Header.Set("Authorization", "token "+res.AccessToken)
		req.Header.Set("Accept", "application/vnd.github.v3+json")
		resp, err = client.Do(req)

		check(err, w, r, user)

		body, err = ioutil.ReadAll(resp.Body)
		check(err, w, r, user)
		defer resp.Body.Close()

		log.Println(string(body))
		//read the response into userdata struct
		err = json.Unmarshal(body, &GitUserEmails)
		check(err, w, r, user)

		userdata.Name = GitUserData.Login
		userdata.Email = GitUserEmails[0].Email

	}

	log.Println("OAuth name:", userdata.Name)
	log.Println("OAuth e-mail:", userdata.Email)

	dbuser := database.GetUserByUserName(userdata.Name).Username
	dbemail := database.GetUserByUserName(userdata.Email).Email

	login(dbuser, dbemail, w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func check(err error, w http.ResponseWriter, r *http.Request, user database.User) {
	var pagedata database.PageData
	pagedata.User = user
	tmpl := template.Must(template.ParseGlob("assets/templates/*.html"))
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "err500", pagedata)
		return
	}

}

var counter int

// check if user exists and login or register new user. change username by adding a number
// if user by this name already exists
func login(name, email string, w http.ResponseWriter, r *http.Request) {
	user := database.GetUserByUserName(email)

	if email == userdata.Email {
		//log user in
		database.AddSession(w, r, user)
		return
	}

	//if there is no user by this e-mail and username
	if email != userdata.Email && name != userdata.Name {
		log.Println("I am here!")
		log.Println(userdata.Name)
		log.Println(userdata.Email)
		//create user to database and log in
		statement, err := database.Db.Prepare("INSERT INTO user (username, password,  email, reg_datetime, role_id) VALUES (?, ?, ?, ?, ?);")
		if err != nil {
			log.Println("Database error, cannot insert new user.")
		}
		defer statement.Close()

		rand.Seed(time.Now().UnixNano())
		//generate random password
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(randSeq(12)), 10)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal server error!")
			return
		}

		statement.Exec(userdata.Name, encryptedPassword, userdata.Email, time.Now(), 1)
		user = database.GetUserByUserName(userdata.Email)
		database.AddSession(w, r, user)
		return
	}

	if counter > 9 {
		userdata.Name = userdata.Name + "0"
		counter = 0
		name = database.GetUserByUserName(userdata.Name).Username
		login(name, email, w, r)
	}

	//if e-mail doesn't exist in database but someone already has a username by the same name
	if email != userdata.Email && name == userdata.Name {
		//change userdame by adding a number to the end
		//check if last char of username is integer
		val, err := strconv.Atoi(string(name[len(name)-1]))

		//if it is not an int, append 1 to username
		if err != nil {
			name = name + "1"
		}

		//if it is an int, make the number higher
		userdata.Name = userdata.Name[:len(name)-1] + strconv.Itoa(val+1)
		counter++
		name = database.GetUserByUserName(userdata.Name).Username
		login(name, email, w, r)
	}
}

//random password generator
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
