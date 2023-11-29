package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var testDB *sql.DB

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		// testint.MはFatal系メソッドを持たないので
		os.Exit(1)
	}

	m.Run() // テストパッケージ内のすべてのユニットテストを実行

	tearDown()
}

func setupTestData() error {
	// cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setupDB.sql") ローカルのmysql経由で実行（本書通り）

	// 以下,自分の環境用に改変
	// dockerのmysqlコンテナ経由で実行(リダイレクトはシェルの機能でありexec.Commandでは使えない)
	setupSql, err := os.Open("./testdata/setupDB.sql")
	if err != nil {
		return err
	}
	cmd := exec.Command("docker", "exec", "-i", "db-for-go", "mysql", "-udocker", "-pdocker", "sampledb")
	// リダイレクト
	cmd.Stdin = setupSql

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {

	// cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/cleanupDB.sql") ローカルのmysql経由で実行（本書通り）

	// 以下,自分の環境用に改変
	// dockerのmysqlコンテナ経由で実行(リダイレクトはシェルの機能でありexec.Commandでは使えない)
	cleanupSql, err := os.Open("./testdata/cleanupDB.sql")
	if err != nil {
		return err
	}
	cmd := exec.Command("docker", "exec", "-i", "db-for-go", "mysql", "-udocker", "-pdocker", "sampledb")
	// リダイレクト
	cmd.Stdin = cleanupSql

	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}

	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}

	return nil
}

func tearDown() {
	cleanupDB()
	testDB.Close()
}
