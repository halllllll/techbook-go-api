[APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx)を写経するリポジトリ

# dif
内容差分

## asdf
`asdf`を使っている。そのせいか、`go.mod`および`go.sum`の記述によってvscode上でエラーが発生することがある

```
Command 'gopls.go_get_package' failed: Error: err: exit status 1: stderr: go: finding module for package　<ローカルのパッケージ>
```
`gopls`のエラーだが、どうやら`asdf`のgolangの`pkg`にインストールされたパッケージを見にいったりしている？
こういう場合、`go.mod`と`go.sum`のそれっぽいパッケージの部分を削除し、コードでimportしてるそれっぽいパッケージの部分も削除。vscodeのGolang拡張機能が入っていれば、そのまま保存すれば現在参照できる形で勝手に修正され、エラーは消えた。

## 環境変数
本書では`go run`時にDBのパスワードなどをオプションで渡し、コードで`os.Getenv`で読み取るのだが、自分の環境ではうまくいかず。環境変数としてオプションを渡す方法は情報がなく不明。

これは`joho/godotenv`を使うことにした
```golang
func init() {
	// 本書で書かれていた「実行時に環境変数をオプションで渡す」が動かなかったのでgodotenvで読み込むことにした
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

// var (
// 	dbUser     = os.Getenv("DB_USER")
// 	dbPassword = os.Getenv("DB_PASSWORD")
// 	dbDatabase = os.Getenv("DB_NAME")
// 	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
// )

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
```

## (p120~)3-1. 前準備 〜DB セットアップ

書籍内ではMySQLのDockerコンテナをローカルにある`mysql`コマンド経由で使っている。自分の環境ではローカルにMySQLがないので、コンテナのシェルに入って直接操作している。

```docker-compose.yaml
services:

  mysql:
    platform: linux/x86_64 <- M1 mac用
    ...各種設定
    volumes:
        - db-volume:/var/lib/mysql
        - ./createTable.sql:/docker-entrypoint-initdb.d/createTable.sql <- これを追加した
```

本書どおりにDBからテーブルを確認する。上記の`docker-compose.yaml`の追記は、コンテナボリュームが作成されるときに実行されるので、すでに作っていた場合はコンテナごと削除して再度作り直します

```
docker-compose down -v
docker-compose up
```

別のターミナルからコンテナに入って確認
```
docker exec -it db-for-go /bin/sh
mysql -u docker sampledb -p
```

* `/usr/bin/mysql: /usr/bin/mysql: cannot execute binary file`となる場合は以下を試す
```
docker exec -it db-for-go  mysql -u docker -p
```

## (p218~)4-5. repositoriesパッケージのテストを完成させよう
**p224**の`exec.Command`について、環境の違いで改変。本書の環境だと`mysql`がローカルにある前提なので、

```golang
cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb",
↪ "--password=docker", "-e", "source ./testdata/cleanupDB.sql")
```

のようになっている。自分の環境では`Docker`のMySQLコンテナ上でDBを動かしているが、`exec.Command`みたいに実行しようとしても、シェル環境ではないので、リダイレクトができない
```golang
// これはリダイレクト演算子をつかっているのでエラーになる
cmd := exec.Command("docker", "exec", "-i", "db-for-go", "mysql", "-udocker", "-pdocker", "sampledb", "<", "./dataset/setupData.sql")
```

ということでChatGPTに聞いたら、`cmd.Stdin`を使って標準入力でファイルを読み込ませろとのこと。

```golang
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
```