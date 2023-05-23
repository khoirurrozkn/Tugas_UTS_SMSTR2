package messageModel

import (
	"utsstrukdat/db"
	"utsstrukdat/entity"
)


func FindMessage(user string) *[]entity.FieldMessage {
	data_Message := db.DataMessage.Next

	var respons []entity.FieldMessage
	for data_Message != nil {
		if data_Message.Data.To == user {

			respons = append(respons, data_Message.Data)
		}
		data_Message = data_Message.Next
	}
	return &respons
}

func CreateMessage(penerima string, pengirim string, pesan string){
	dataMessage := &db.DataMessage

	data := &entity.Message{
		Data: entity.FieldMessage{
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