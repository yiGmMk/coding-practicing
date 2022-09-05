#!/usr/bin/env bash

sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
go build -o $sh_dir/main -ldflags "-X main.version=v1.0.1" $sh_dir/main.go 
$sh_dir/main