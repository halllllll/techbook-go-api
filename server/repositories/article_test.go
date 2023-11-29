package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/halllllll/techbook-go-api/server/models"
	"github.com/halllllll/techbook-go-api/server/repositories"
)

func TestSelectArticleList(t *testing.T) {

	expectedNum := 4
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {

	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPoost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  3,
			},
		},
		{
			testTitle: "subtest2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd Post",
				Contents: "Second Blog Post",
				UserName: "saki",
				NiceNum:  9,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("get %d but want %d\n", got.ID, test.expected.ID)
			}

			if got.Title != test.expected.Title {
				t.Errorf("get %s but want %s\n", got.Title, test.expected.Title)
			}

			if got.Contents != test.expected.Contents {
				t.Errorf("get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}

}
