package messageModel

import (
	"utsstrukdat/db"
)


func FindMessage(user string) *[]db.FieldMessage {
	var data_Message *db.Message = &db.DataMessage
	temp := data_Message.Next

	var respons []db.FieldMessage
	for temp != nil {
		if temp.Data.To == user {

			respons = append(respons, temp.Data)
		}
		temp = temp.Next
	}
	return &respons
}
