#!/usr/bin/env bash
###
 # @Author: yiGmMk marvelousme@163.com
 # @Date: 2022-09-05 17:01:41
 # @LastEditors: yiGmMk marvelousme@163.com
 # @LastEditTime: 2022-09-05 17:37:17
 # @FilePath: /go-tool/home/admin/code/coding-practicing/base/build/version copy.sh
 # @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
### 
sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

rm $sh_dir/main
echo "add tag b"
go build -o $sh_dir/main -tags="b" $sh_dir/*.go
$sh_dir/main

rm $sh_dir/main
echo "add tag a"
sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
go build -o $sh_dir/main -tags="a" $sh_dir/*.go
$sh_dir/main

rm $sh_dir/main
echo "add tag a"
sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
go build -o $sh_dir/main -ldflags "-X main.version=v1.0.1" -tags="a" $sh_dir/*.go
$sh_dir/main

rm $sh_dir/main
echo "add tag a,b"
go build -o $sh_dir/main -ldflags "-X main.version=v1.0.2" -tags="a,b" $sh_dir/*.go
$sh_dir/main