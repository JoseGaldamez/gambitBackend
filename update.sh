#!/bin/bash

git add .
git commit -m "update"
git push origin main

env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o . main.go
zip main.zip main
