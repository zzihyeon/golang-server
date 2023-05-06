# golang-server
graphQL, restAPI, postgresql, mongoDB

# Backend with Golang, postgresql
- [Based on this lecture](https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)
## Setting commands
- [Docker postgresql](https://hub.docker.com/_/postgres)
- [Golang migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
### Install
```sh
brew install golang-migrate
brew install sqlc

go mod init github.com/zzihyeon/golang-server.git
go mod tidy # installs go library
```

### Docker commands
```sh
docker ps -a
docker images
docker pull postgres:13-alpine
docker run --name posql -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:13-alpine
docker exec -it posql psql -U root # localhost 는 No 비번
docker logs posql
docker stop posql
docker exec -it posql createdb --username=root --owner=root cavea
docker rm --force posql
```

docker file 사용
```
docker build -t ${생성할 이미지 이름}:${테그} .
```

docker compose 사용
```
cd docker
docker-compose up -d
```

docker log 확인
```
docker logs cavea-dev -f --tail=100
```

### Migrate
#### create migration
```sh
migrate create -ext sql -dir db/migration -seq init_schema
```
#### exec migration
```sh
migrate -path=/root/dev/goback/db/migration -database=${dbpath} up 1

db가 정상적으로 동작해야 유효함
migrate -path ./db/migration -database ${dbpath} up 1
migrate -path ./db/migration -database ${dbpath} down 1
```

### SQLC
- [SQLC Github](https://github.com/kyleconroy/sqlc)
  
```sh
sqlc version
sqlc init # yaml file
sqlc generate
```

### GraphQL
- [gqlgen Github](https://github.com/99designs/gqlgen)
#### 아래의 command로 .graphqls에 수정한 사항을 반영할 수 있다.
```sh
go run github.com/99designs/gqlgen
```

### POSTGRESQL
- docker 설정
```
docker pull postgres:13-alpine
docker run --name posql -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:13-alpine
docker exec -it posql psql -U root
```
- create user
```
CREATE ROLE test WITH LOGIN PASSWORD 'test';
\du
ALTER ROLE test CREATEDB;
\du
/c test
CREATE DATABASE test;
\list
```