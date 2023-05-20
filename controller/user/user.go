package userController

import (
	"utsstrukdat/db"
	"utsstrukdat/model/user"
	"utsstrukdat/config/crypt"
)

func Register(req db.FieldUser, verPassword string) int {

	check := userModel.FindOne(req.Username)

	if check.Username != "" {
		return 409
	}
	
	if req.Password != verPassword {
		return 400
	}

	req.Password = crypt.Encode(req.Password)

	userModel.Create(req)
	
	return 200
	
}

func Login(req db.FieldUser) db.FieldUser {

	check := userModel.FindOne(req.Username)
	
	if check.Username == req.Username && crypt.Decode(check.Password) == req.Password {
		return check
	}
	check.Password = crypt.Encode(check.Password)
	return db.FieldUser{}
}

func ShowPostByAccount(username string) *[]db.FieldPost {

	check := userModel.FindOne(username)

	if check.Username == "" {
		return nil
	}

	response := userModel.FindUserAndPost(username)

	return response
}

func SearchAccount(username string) string {
	check := userModel.FindOne(username)

	if check.Username != ""{
		return check.Username
	}

	return ""
}
