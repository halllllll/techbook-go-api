package testdata

import "gihtub.com/halllllll/techbook-go-api/server/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  4,
	},
	models.Article{
		ID:       2,
		Title:    "2nd Post",
		Contents: "Second Blog Post",
		UserName: "saki",
		NiceNum:  9,
	},
}
