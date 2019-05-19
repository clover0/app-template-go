# README
* Go version 1.11
* web framework: echo,
* ORM(sql):sqlx, http://jmoiron.github.io/sqlx/
* SessionStore: Redis

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
docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/<FILE_NAME> -U postgres -d auth465


## Testing
* set up database  
`docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/<FILE_NAME> -U postgres -d auth465_test`

* run all test   
`docker exec -it auth465 go test -v ./...`  
