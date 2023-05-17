package messageModel

import (
	"utsstrukdat/db"
)

func FindMessage(user string) *[]db.FieldMessage {
	var dataMessage *db.Message = &db.DataMessage
	temp := dataMessage.Next

	var response []db.FieldMessage
	for temp != nil {
		if temp.Data.To == user {

			response = append(response, temp.Data)
		}
		temp = temp.Next
	}
	return &response
}

func CreateMessage(penerima string, pengirim string, pesan string){
	var dataMessage *db.Message = &db.DataMessage

	data := &db.Message{
		Data: db.FieldMessage{
			From: pengirim,
			To: penerima,
			Message: pesan,
		},
	}

	if dataMessage.Next == nil {
		dataMessage.Next = data
	}else{
		data.Next = dataMessage.Next
		dataMessage.Next = data
	}
}