#!/usr/bin/env bash
mkdir -p goreport
go test -v -json -cover -coverprofile goreport/cover.out ./... >goreport/report.jsonl
go tool cover -html=goreport/cover.out -o goreport/index.html