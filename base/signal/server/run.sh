#!/usr/bin bash

script_file_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
go build -o $script_file_dir/main $script_file_dir/gracefully-exit-with-notify.go
$script_file_dir/main