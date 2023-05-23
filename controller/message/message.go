package messageController

import (
	"utsstrukdat/model/message"
	"utsstrukdat/model/user"
	"utsstrukdat/entity"
)

func ShowMessage(user string) *[]entity.FieldMessage {

	response := messageModel.FindMessage(user)

	return response
}

func SendMessage(penerima string, pengirim string, pesan string) int {
	
	check := userModel.FindOne(penerima)

	if check.Username == "" {
		return 404
	}

	messageModel.CreateMessage(penerima, pengirim, pesan)
	return 200
}