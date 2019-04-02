# A
* Go version 1.11
* web framework: echo,
* ORM(sql):sqlx, http://jmoiron.github.io/sqlx/

## setup
### database
* データベースの作成
`docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/0.sql -U postgres`

* マイグレーション実行
docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/ファイル名 -U postgres