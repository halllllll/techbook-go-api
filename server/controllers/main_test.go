package controllers_test

import (
	"testing"

	"gihtub.com/halllllll/techbook-go-api/server/controllers"
	"gihtub.com/halllllll/techbook-go-api/server/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	// dbUser := "docker"
	// dbPassword := "docker"
	// dbDatabase := "sampledb"
	// dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// db, err := sql.Open("mysql", dbConn)

	// if err != nil {
	// 	log.Println("db setup fail")
	// 	os.Exit(1)
	// }
	// defer db.Close()

	// ser := services.NewMyAppService(db)
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
