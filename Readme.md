# Go restful test

1 CRUD endpoint involving GORM, mysql, fiber, migration & seeder, swaggo, docker

## Prerequisites

Minimum system requirements

```
Go 1.19
Mysql 8~
```

## Installing & Running

Copy .env-example to .env and fill the required configuration

```
go mod tidy
go run main.go
```

## Running the tests

In order to run entire test, user should have & granted docker usage

```
go test ./...
```


##FAQ:

** Q: Docker test?
** A: Yes, bcs it's pain in the ass trying to mock up all interface, when testing database / environment related system it's better to go all out using integrated test instead

** Q: Why there's no api versioning?
** A: Api versioning could be ran as separate instance with help of api gateway management like traefik or kong, as development goes on, it will be hindrance to handle backward compability in same upstream.

