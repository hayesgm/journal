#! /bin/bash

go build journal.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o journal.linux journal.go
