package userModel

import (
	"utsstrukdat/db"
)

func FindOne(username string) db.FieldUser {
	var userDB *db.User = &db.DataUser
	current := userDB.Next

	for current != nil {
		if current.Data.Username == username {
			return current.Data
		}
		current = current.Next
	}

	return db.FieldUser{}
}

func Create(req db.FieldUser) {
	var userDB *db.User = &db.DataUser

	newUser := &db.User{
		Data: db.FieldUser{
			Username: req.Username,
			Password: req.Password,
		},
	}

	if userDB.Next == nil {
		userDB.Next = newUser
	} else {
		newUser.Next = userDB.Next
		userDB.Next = newUser
	}
}

func FindUserAndPost(username string) *[]db.FieldPost {
	var postDB *db.Post = &db.DataPost
	current := postDB.Next

	var response []db.FieldPost

	for current != nil {
		if current.Data.Author == username {
			response = append(response, current.Data)
		}
		current = current.Next
	}

	return &response
}
