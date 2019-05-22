# Auth-Server-Template
Auth-Server-Template is a template to build authentication server quickly with Go.
This project shows you how to make it by DI with Go, the mechanism of authentication, 
and the architecture of the web application.

# Env
* Go version 1.11
* web framework: Echo,
* ORM(sql):sqlx, http://jmoiron.github.io/sqlx/
* SessionStore: Redis4.0

## For Developer
* DI code gen
`wire`  

* run app  
`docker-compose up`

* stop app  
`docker-compose stop`

* reset containers  
`docker-compose down`  
This command deletes database and session store!

### check api

* do sign_in  
`curl -X POST -H "Content-Type: application/json" -d '{"email":"<Email!>", "password":"<Password!>"}' http://localhost:1323/a/auth -c cookie.txt`  

* do sign_out  
`curl -X POST -b cookie.txt http://localhost:1323/a/sign_out -c cookie.txt`  


## How to Setup
### Database
* create databse  
`docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/0.sql -U postgres`

* execute migrate
`docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/<FILE_NAME> -U postgres -d auth465`


## For Testing
* set up database  
`docker exec -it postgresql_auth465 psql -f /home/db/migration/sql/<FILE_NAME> -U postgres -d auth465_test`

* run all test   
`docker exec -it auth465 go test -v ./...`  
