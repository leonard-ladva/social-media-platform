package request

import (
	"forum-test/database"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateDummyData() error {
	//create admin user
	username := "admin"
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), 10)
	if err != nil {
		return err
	}
	email := "admin@test.com"

	statement, err := database.Db.Prepare("INSERT INTO user (username, password, email, reg_datetime, role_id) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Println("Database error, cannot insert new user.")
	}
	defer statement.Close()
	statement.Exec(username, encryptedPassword, email, time.Now(), 3)

	//create moderator user
	username = "moderator"
	encryptedPassword, err = bcrypt.GenerateFromPassword([]byte("moderator"), 10)
	if err != nil {
		return err
	}
	email = "moderator@test.com"

	statement, err = database.Db.Prepare("INSERT INTO user (username, password, email, reg_datetime, role_id) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Println("Database error, cannot insert new user.")
	}
	defer statement.Close()
	statement.Exec(username, encryptedPassword, email, time.Now(), 2)

	//create 3 dummy users
	for i := 1; i < 4; i++ {
		username := "test" + strconv.Itoa(i)
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte("test"+strconv.Itoa(i)), 10)
		if err != nil {
			return err
		}
		email := "test" + strconv.Itoa(i) + "@test.com"

		statement, err := database.Db.Prepare("INSERT INTO user (username, password, email, reg_datetime, role_id) VALUES (?, ?, ?, ?, ?);")
		if err != nil {
			log.Println("Database error, cannot insert new user.")
		}
		defer statement.Close()
		statement.Exec(username, encryptedPassword, email, time.Now(), 1)
	}

	//create user roles
	statement, err = database.Db.Prepare("INSERT INTO role (role_id, role) VALUES (0, 'guest'), (1, 'user'), (2, 'moderator'), (3, 'administrator');")
	if err != nil {
		log.Println("Database error, cannot insert roles.")
	}
	defer statement.Close()
	statement.Exec()

	//create categories
	content, err := ioutil.ReadFile("assets/conf_categories.txt")
	if err != nil {
		log.Fatal(err)
	}
	categories := strings.Split(string(content), "\n")
	for i := range categories {
		category := strings.Split(string(categories[i]), ";")
		statement, err := database.Db.Prepare("INSERT INTO category (title, description, img_link) VALUES (?, ?, ?);")
		if err != nil {
			log.Println("Database error, cannot insert new category.")
		}
		defer statement.Close()
		statement.Exec(category[0], category[1], category[2])
	}

	//create posts
	content, err = ioutil.ReadFile("assets/conf_posts.txt")
	if err != nil {
		log.Fatal(err)
	}
	posts := strings.Split(string(content), "\n")
	for i := range posts {
		post := strings.Split(string(posts[i]), ";")
		statement, err := database.Db.Prepare("INSERT INTO post (title, content, user_id, datetime, status) VALUES (?, ?, ?, ?, ?);")
		if err != nil {
			log.Println("Database error, cannot insert new post.")
		}
		defer statement.Close()
		statement.Exec(post[0], post[1], post[2], time.Now(), post[3])
	}

	//create post_category
	for postnr := 1; postnr <= 5; postnr++ {
		skiplist := []int{}
		for j := 0; j < 2; j++ {
			rand.Seed(time.Now().UnixNano())
			category_id := rand.Intn(5) + 1
			if contains(skiplist, category_id) {
				continue
			}
			statement, err := database.Db.Prepare("INSERT INTO post_category (post_id, category_id) VALUES (?, ?);")
			if err != nil {
				log.Println("Database error, cannot insert new post_category.")
			}
			defer statement.Close()
			statement.Exec(postnr, category_id)
			skiplist = append(skiplist, category_id)
		}
	}

	//create post_reaction
	for postnr := 1; postnr <= 5; postnr++ {
		skiplist := []int{}
		for react := 0; react < 6; react++ {
			rand.Seed(time.Now().UnixNano())
			x := (rand.Intn(5) + 1)
			if contains(skiplist, x) {
				continue
			}
			reaction := "like"
			if (postnr+x)%2 == 1 {
				reaction = "dislike"
			}
			statement, err := database.Db.Prepare("INSERT INTO post_reaction (post_id, user_id, type, datetime) VALUES (?, ?, ?, ?);")
			if err != nil {
				log.Println("Database error, cannot insert new post_reaction")
			}
			defer statement.Close()
			statement.Exec(postnr, x, reaction, time.Now())
			skiplist = append(skiplist, x)
		}
	}

	//create comments
	content, err = ioutil.ReadFile("assets/conf_comments.txt")
	if err != nil {
		log.Fatal(err)
	}
	comments := strings.Split(string(content), "\n")
	for i := range comments {
		comm := strings.Split(string(comments[i]), ";")
		statement, err := database.Db.Prepare("INSERT INTO comment (post_id, user_id, content, datetime) VALUES (?, ?, ?, ?);")
		if err != nil {
			log.Println("Database error, cannot insert new comment.")
		}
		defer statement.Close()
		statement.Exec(comm[0], comm[1], comm[2], time.Now())
	}

	//create comment_reaction
	for commnr := 1; commnr <= 10; commnr++ {
		skiplist := []int{}
		for react := 0; react < 6; react++ {
			rand.Seed(time.Now().UnixNano())
			x := (rand.Intn(5) + 1)
			if contains(skiplist, x) {
				continue
			}
			reaction := "like"
			if (commnr+x)%2 == 1 {
				reaction = "dislike"
			}
			statement, err := database.Db.Prepare("INSERT INTO comment_reaction (comment_id, user_id, type, datetime) VALUES (?, ?, ?, ?);")
			if err != nil {
				log.Println("Database error, cannot insert new post_reaction")
			}
			defer statement.Close()
			statement.Exec(commnr, x, reaction, time.Now())
			skiplist = append(skiplist, x)
		}
	}
	log.Println("Auditing tip: for your convenience, we added categories, posts, comments and reactions")
	log.Println("Also, we created the following users: admin, moderator, test1, test2, test3. Password for every generated user is the same as username")
	return nil
}

func contains(list []int, inp int) bool {
	for i := range list {
		if list[i] == inp {
			return true
		}
	}
	return false
}
