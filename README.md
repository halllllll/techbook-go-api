[APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx)を写経するリポジトリ

## dif
### (p120~)3-1. 前準備 〜DB セットアップ

ローカルに`mysql`がない用のメモ。

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

