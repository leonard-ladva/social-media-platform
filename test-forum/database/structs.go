package database

// Account roles
const GUEST = 0
const USER = 1
const MODERATOR = 2
const ADMIN = 3

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Date     string
	RoleID   int
	Role     string
	Unread   int
}

type Post struct {
	ID           int
	Title        string
	Content      string
	Image        string
	Date         string
	Author       string
	Likes        int
	Dislikes     int
	Comments     int
	Categories   []string
	Reaction     string
	Status       string
	UserComments []Comment
}

type Category struct {
	ID          int
	Title       string
	Description string
	Image       string
	Checked     string
}

type Comment struct {
	ID       int
	Post_ID  int
	Content  string
	Date     string
	Author   string
	Likes    int
	Dislikes int
	Reaction string
}

type Session struct {
	Session_ID int
	Datetime   string
	User_ID    int
	Uuid       string
}

type PageData struct {
	User          User
	Message       Message
	Posts         []Post
	Categories    []Category
	Post          Post
	Comments      []Comment
	Notifications []Notification
	Comment       Comment
	Authdata      Authdata
	CALLBACK_URI  string
}

type Message struct {
	Msg1 string
	Msg2 string
	Msg3 string
	Msg4 string
}

type Likes struct {
	Likes []Like
}

type Like struct {
	PostID int
	Type   string
}

type Notification struct {
	ID       int
	User_ID  int
	Content  string
	Link     string
	Priority int
	Date     string
	Read     int
}

type Authdata struct {
	GoogleClientID   string
	FacebookClientID string
	GitHubClientID   string
}
