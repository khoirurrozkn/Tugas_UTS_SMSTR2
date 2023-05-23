package userController

import (
	"utsstrukdat/model/user"
	"utsstrukdat/config/crypt"
	"utsstrukdat/entity"
)

func Register(req entity.FieldUser, verPassword string) int {

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

func Login(req entity.FieldUser) entity.FieldUser {

	check := userModel.FindOne(req.Username)
	
	if check.Username == req.Username && crypt.Decode(check.Password) == req.Password {
		return check
	}
	check.Password = crypt.Encode(check.Password)
	return entity.FieldUser{}
}

func ShowPostByAccount(username string) *[]entity.FieldPost {

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
