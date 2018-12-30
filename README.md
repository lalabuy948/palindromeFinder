# Palindrome finder

> SSorting 35k words in 0.01 sec

## How to run

```sh
go run main.go

go build main

./main
```

## Time

JIT

```sh
➜ time go run main.go
ротатор
go run main.go  0.21s user 0.10s system 130% cpu 0.235 total
```

Compiled

```sh
➜ time ./main
ротатор
./main  0.01s user 0.00s system 86% cpu 0.019 total
```
