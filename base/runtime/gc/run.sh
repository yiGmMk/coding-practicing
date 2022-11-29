#!/usr/bin/env bash

dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

go build -o $dir/main $dir/gc.go

GODEBUG='gctrace=1' GOGC=30 $dir/main