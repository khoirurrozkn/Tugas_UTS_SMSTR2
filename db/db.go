package db

type FieldUser struct {
	Username string
	Password string
}

type User struct {
	Data FieldUser
	Next *User
}

type FieldPost struct {
	Author string
	Category string
	Title string
	Body string
}

type Post struct {
	Data FieldPost
	Next *Post
}

type FieldMessage struct {
	From string
	To string
	Message string
}

type Message struct {
	Data FieldMessage
	Next *Message
}

var DataUser User
var DataPost Post
var DataMessage Message