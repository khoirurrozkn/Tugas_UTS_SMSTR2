package userModel

import (
	"utsstrukdat/db"
)

func FindOne(username string) db.FieldUser {
	var dataUser *db.User = &db.DataUser
	temp := dataUser.Next

	for temp != nil {
		if temp.Data.Username == username {

			return temp.Data
		}
		temp = temp.Next
	}

	return db.FieldUser{}
}

func Create(req db.FieldUser){
	var dataUser *db.User = &db.DataUser

	data := &db.User{
		Data: db.FieldUser{
			Username: req.Username,
			Password: req.Password,
		},
	}

	if dataUser.Next == nil {
		dataUser.Next = data
	}else{
		data.Next = dataUser.Next
		dataUser.Next = data
	}
}

func FindUserAndPost(username string) *[]db.FieldPost {
	var dataPost *db.Post = &db.DataPost
	temp := dataPost.Next

	var response []db.FieldPost

	for temp != nil {
		if temp.Data.Author == username {
			
			response = append(response, temp.Data)
		}
		temp = temp.Next
	}
	return &response
}