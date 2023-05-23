package postController

import (
	"utsstrukdat/model/post"
	"utsstrukdat/entity"
)
func ShowPost() *[]entity.FieldPost {
	
	return postModel.Find()
}

func InsertPost(req *entity.FieldPost) int {
	check := postModel.FindOne(req.Title)

	if check != nil {
		return 409
	}

	postModel.Create(req)
	return 200
}

func UpdatePost(title string, body string, author string) int {
	check := postModel.FindOne(title)

	if check.Title != "" {
		if check.Author != author {
			return 403
		}

		postModel.FindByTitleAndUpdate(title, body)
		return 200
	}

	return 404
}

func DeletePost(title string, author string) int {
	check := postModel.FindOne(title)

	if check.Title != "" {
		if check.Author != author {
			return 403
		}

		postModel.FindByTitleAndDelete(title)
		return 200
	}

	return 404
}

func ShowByCategory(category string) *[]entity.FieldPost {

	response := postModel.FindPostByCategory(category)

	return response
}