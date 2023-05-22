package postModel

import (
	"utsstrukdat/db"
)

func Find() *[]db.FieldPost {
	var isipost *db.Post = &db.DataPost
	temp := isipost.Next

	var data []db.FieldPost

	for temp != nil {
		data = append(data, temp.Data)
		temp = temp.Next
	}

	return &data
}

func FindOne(judul string) *db.FieldPost {
	var isipost *db.Post = &db.DataPost
	temp := isipost.Next

	for temp != nil {
		if temp.Data.Title == judul {

			return &temp.Data
		}
		temp = temp.Next
	}
	
	return nil
}

func Create(req *db.FieldPost){
	var isipost *db.Post = &db.DataPost

	data := &db.Post{
		Data: db.FieldPost{
			Author: req.Author,
			Category: req.Category,
			Title: req.Title,
			Body: req.Body,
		},
	}

	if isipost.Next == nil {
		isipost.Next = data
	}else {
		data.Next = isipost.Next
		isipost.Next = data
	}
}

func FindByTitleAndUpdate(title string, body string){
	var isipost *db.Post = &db.DataPost
	temp := isipost.Next

	for temp != nil {
		if temp.Data.Title == title {

			temp.Data.Body = body
			break
		}
		temp = temp.Next
	}
}

func FindByTitleAndDelete(title string){
	var isipost *db.Post = &db.DataPost

	after := isipost.Next
	before := isipost

	for after != nil {
		if after.Data.Title == title {
			if  after == isipost.Next {
				isipost.Next = after.Next
			}else{
				before.Next = after.Next
			}
			break
		}
		before = after
		after = after.Next
	}
}

func FindPostByCategory(category string) *[]db.FieldPost {
	var isipost *db.Post = &db.DataPost
	temp := isipost.Next

	var response []db.FieldPost

	for temp != nil {
		if temp.Data.Category == category {

			response = append(response, temp.Data)
		}
		temp = temp.Next
	}

	return &response
}
