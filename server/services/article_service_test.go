package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"gihtub.com/halllllll/techbook-go-api/server/services"

	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbName := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)

	db, err := sql.Open("mysql", dbConn)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewMyAppService(db)

	// 個別のベンチマークの測定

	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	// ここから先の時間を計測する
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
