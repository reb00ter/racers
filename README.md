# Race Genteelest

Vote for more accurate racers

# Stack

## Backend

- HTTP Middleware [Echo](https://echo.labstack.com/)
- ORM library [gorm](https://github.com/jinzhu/gorm)
- Configuration [env](https://github.com/caarlos0/env)
- Load ENV variables from .env file [godotenv](https://github.com/joho/godotenv)
- Payload validation [validator](https://github.com/go-playground/validator)
- Cache [Redis](https://github.com/go-redis/redis)
- Localization [gotext](https://github.com/leonelquinteros/gotext)
- Database [CockroachDB](https://github.com/cockroachdb/cockroach)
- Releasing [goreleaser](https://github.com/goreleaser/goreleaser)

## Frontend

- Server side templating [Go Templates](https://golang.org/pkg/text/template/)
- Module Bundler [Parcel bundler](https://github.com/parcel-bundler/parcel)
- Javascript UI library [React](https://github.com/facebook/react)

# Getting Started

## Project structure

Follows https://github.com/golang-standards/project-layout

## Building From Source

This project requires Go +1.13 and Go module support.

To build the project run:

```
make
```

## Bootstrap infrastructure and run application

This project requires docker and docker compose to run the required services.

1. To run the services:

```
docker-compose up
```

2. To create database

```
docker run --network="host" -it cockroachdb/cockroach:v19.2.1 sql --insecure -e "$(cat ./scripts/create.db.sql)"
```

3. Build [web application](web/README.md)

4. Start server

```
go run main.go
```

5. Navigate to users list [page](http://127.0.0.1/users)

## CI and Static Analysis

### CI

All pull requests will run through CI, which is currently hosted by Github-CI.
Community contributors should be able to see the outcome of this process by looking at the checks on their PR.
Please fix any issues to ensure a prompt review from members of the team.

### Static Analysis

This project uses the following static analysis tools.
Failure during the running of any of these tools results in a failed build.
Generally, code must be adjusted to satisfy these tools, though there are exceptions.

- [go vet](https://golang.org/cmd/vet/) checks for Go code that should be considered incorrect.
- [go fmt](https://golang.org/cmd/gofmt/) checks that Go code is correctly formatted.
- [golangci-lint](https://github.com/golangci/golangci-lintt) checks for things like: unused code, code that can be simplified, code that is incorrect and code that will have performance issues.
- [go mod tidy](https://tip.golang.org/cmd/go/#hdr-Add_missing_and_remove_unused_modules) ensures that the source code and go.mod agree.

# Releasing

When a new tag is pushed, the version is released with [goreleaser](https://github.com/goreleaser/goreleaser).

```
$ git tag -a v0.1.0 -m "First release"
$ git push origin v0.1.0 # => want to release v0.1.0
```

# Tooling

- IDE plugin [vscode-go](https://github.com/Microsoft/vscode-go)
- Administration of cockroachdb [DBeaver](https://dbeaver.io/)
- REST client [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop?hl=en)
- Go testing in the browser [go-convey](https://github.com/smartystreets/goconvey)
- Benchmarking [bombardier](http://github.com/codesenberg/bombardier)

# Documentation

```
$ godoc github.com/starptech/go-web/pkg/controller
$ godoc -http=:6060
```

Visit localhost:6060 and search for `go-web`

# Benchmarking

```
$ bombardier -c 10 -n 10000 http://localhost:8080/users
```

# Cockroachdb Cluster overview

http://localhost:8111/

## Deploy on Heroku

[![Heroku Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/StarpTech/go-web)

# Further reading

- http://www.alexedwards.net/blog/organising-database-access
- https://12factor.net/
- https://dev.otto.de/2015/09/30/on-monoliths-and-microservices/
