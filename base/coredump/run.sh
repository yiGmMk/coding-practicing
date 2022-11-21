#!/usr/bin bash

ulimit -c unlimited
go build main.go

GOTRACEBACK=crash ./main