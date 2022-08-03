package models

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// mockup DB.
var articleList = []article{
	article{1, "Hello world!", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"},
	article{2, "Hello Go!", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"},
	article{3, "Hello Gin!", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt"},
}

func GetAllArticles() []article {
	return articleList
}

//* (*article, error) : return type
func GetArticleByID(id int) (*article, error) {
	//* _ index, a item
	for _, item := range articleList {
		if item.ID == id {
			return &item, nil // return data, no error
		}
	}
	return nil, errors.New("Article for this id is not found!")
}
