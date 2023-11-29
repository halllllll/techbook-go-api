package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		// testint.MはFatal系メソッドを持たないので
		os.Exit(1)
	}

	m.Run() // テストパッケージ内のすべてのユニットテストを実行

	tearDown()
}

func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}

	return nil
}

func tearDown() {
	testDB.Close()
}
