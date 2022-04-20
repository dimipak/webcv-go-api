#!/bin/bash

cd /home/ubuntu/admin/base

go get ./...

env GOOS=linux GOARCH=amd64 go build -o /home/ubuntu/admin/admin main.go