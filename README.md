# Got

Reimplementing git from scratch, in go

## Hello world

Run go in docker

```sh
docker compose build dev
docker compose run dev
```

then, for example:

```sh
go fmt
go build
./got
➜ Hello world!
```

## create a `got` alias

```sh
docker compose build got
alias got="docker compose run got"
got
➜ Hello world!
```
