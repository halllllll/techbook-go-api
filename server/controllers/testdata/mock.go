package testdata

import "gihtub.com/halllllll/techbook-go-api/server/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMock) GetArticleService(articleID int) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return articleTestData, nil
}

func (s *serviceMock) PostNiceService(models.Article) (models.Article, error) {
	return articleTestData[0], nil
}
