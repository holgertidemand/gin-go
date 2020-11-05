package main

import (
	"errors"
)

//Article - a
type Article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

var articlesList = []Article{
	Article{
		ID: 1,
		Title: "Article 1",
		Content: "Article 1 body",
	},
	Article{
		ID: 2,
		Title: "Article 2",
		Content: "Article 2 body",
	},
}


func getAllArticles() []Article{
	return articlesList
}

func getArticleByID(id int) (*Article, error) {

	for _, article := range articlesList {

		if article.ID == id {
			return &article, nil
		 
		}


	}

	return nil, errors.New("Missing Id") 
 


}