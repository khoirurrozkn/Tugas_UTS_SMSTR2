package userModel

import (
	"utsstrukdat/db"
	"utsstrukdat/entity"
)

func FindOne(username string) entity.FieldUser {
	userDB := db.DataUser.Next

	for userDB != nil {
		if userDB.Data.Username == username {
			return userDB.Data
		}
		userDB = userDB.Next
	}

	return entity.FieldUser{}
}

func Create(req entity.FieldUser) {
	userDB := &db.DataUser

	newUser := &entity.User{
		Data: entity.FieldUser{
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

func FindUserAndPost(username string) *[]entity.FieldPost {
	postDB := db.DataPost.Next

	var response []entity.FieldPost

	for postDB != nil {
		if postDB.Data.Author == username {
			response = append(response, postDB.Data)
		}
		postDB = postDB.Next
	}

	return &response
}
