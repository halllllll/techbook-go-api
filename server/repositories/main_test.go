package repositories_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var testDB *sql.DB

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
}

func setupTestData() error {
	setupSql, err := os.Open("./testdata/setupDB.sql")
	if err != nil {
		return err
	}
	cmd := exec.Command("docker", "exec", "-i", "db-for-go", "mysql", "-udocker", "-pdocker", "sampledb")
	// リダイレクト
	cmd.Stdin = setupSql

	err = cmd.Run()
	return err
}

func cleanupDB() error {
	setupSql, err := os.Open("./testdata/cleanupDB.sql")
	if err != nil {
		return err
	}
	cmd := exec.Command("docker", "exec", "-i", "db-for-go", "mysql", "-udocker", "-pdocker", "sampledb")
	// リダイレクト
	cmd.Stdin = setupSql

	err = cmd.Run()
	return err
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	// パッケージに含まれるすべてのテストコードを実行
	m.Run()

	teardown()
}

func connectDB() error {
	var (
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
	)

	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}

	return err

}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}

	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup")
		return err
	}

	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}

	return nil
}

func teardown() {
	cleanupDB()
	testDB.Close()
}
