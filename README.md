# README
* Go version 1.11
* web framework: echo,
* ORM(sql):sqlx, http://jmoiron.github.io/sqlx/



## Develop
* DI code gen
`wire`  

* run app  
`docker-compose up`

* stop app  
`docker-compose stop`

* reset containers  
`docker-compose down`  
Delete database, session store!

## Setup
### Database
* create databse  
`docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/0.sql -U postgres`

* execute migrate
docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/ファイル名 -U postgres -d auth465