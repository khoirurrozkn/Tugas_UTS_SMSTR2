package postModel

import (
	"utsstrukdat/db"
	"utsstrukdat/entity"
)

func Find() *[]entity.FieldPost {
	isipost := db.DataPost.Next

	var data []entity.FieldPost

	for isipost != nil {
		data = append(data, isipost.Data)
		isipost = isipost.Next
	}

	return &data
}

func FindOne(judul string) *entity.FieldPost {
	isipost := db.DataPost.Next

	for isipost != nil {
		if isipost.Data.Title == judul {

			return &isipost.Data
		}
		isipost = isipost.Next
	}
	
	return nil
}

func Create(req *entity.FieldPost){
	isipost := &db.DataPost

	data := &entity.Post{
		Data: *req,
	}

	if isipost.Next == nil {
		isipost.Next = data
	}else {
		data.Next = isipost.Next
		isipost.Next = data
	}
}

func FindByTitleAndUpdate(title string, body string){
	isipost := &db.DataPost

	for isipost.Next != nil {
		if isipost.Data.Title == title {

			isipost.Data.Body = body
			break
		}
		isipost = isipost.Next
	}
}

func FindByTitleAndDelete(title string){
	isipost := &db.DataPost

	for isipost.Next != nil {
		if isipost.Next.Data.Title == title {
			isipost.Next = isipost.Next.Next
			break
		}
		isipost = isipost.Next
	}
}

func FindPostByCategory(category string) *[]entity.FieldPost {
	isipost := db.DataPost.Next

	var response []entity.FieldPost

	for isipost != nil {
		if isipost.Data.Category == category {

			response = append(response, isipost.Data)
		}
		isipost = isipost.Next
	}

	return &response
}
