package postModel

import (
	"utsstrukdat/db"
)

func Find() *[]db.FieldPost {
	var dataPost *db.Post = &db.DataPost
	temp := dataPost.Next

	var data []db.FieldPost

	for temp != nil {
		data = append(data, temp.Data)
		temp = temp.Next
	}

	return &data
}

func FindOne(judul string) *db.FieldPost {
	var dataPost *db.Post = &db.DataPost
	temp := dataPost.Next

	for temp != nil {
		if temp.Data.Title == judul {

			return &temp.Data
		}
		temp = temp.Next
	}
	
	return nil
}

func Create(req *db.FieldPost){
	var dataPost *db.Post = &db.DataPost

	data := &db.Post{
		Data: db.FieldPost{
			Author: req.Author,
			Category: req.Category,
			Title: req.Title,
			Body: req.Body,
		},
	}

	if dataPost.Next == nil {
		dataPost.Next = data
	}else {
		data.Next = dataPost.Next
		dataPost.Next = data
	}
}

func FindByTitleAndUpdate(title string, body string){
	var dataPost *db.Post = &db.DataPost
	temp := dataPost.Next

	for temp != nil {
		if temp.Data.Title == title {

			temp.Data.Body = body
			break
		}
		temp = temp.Next
	}
}

func FindByTitleAndDelete(title string){
	var dataPost *db.Post = &db.DataPost

	after := dataPost.Next
	before := dataPost

	for after != nil {
		if after.Data.Title == title {
			if  after == dataPost.Next {
				dataPost.Next = after.Next
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
	var dataPost *db.Post = &db.DataPost
	temp := dataPost.Next

	var response []db.FieldPost

	for temp != nil {
		if temp.Data.Category == category {

			response = append(response, temp.Data)
		}
		temp = temp.Next
	}

	return &response
}