#!/usr/bin/bash 

sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
go build -o $sh_dir/main $sh_dir/main.go 
go tool nm $sh_dir/main | grep "main\."




#   4ba720 R main..dict.Sort[float64]
#   4ba730 R main..dict.Sort[int]
#   516760 D main..inittask                             // 包依赖初始化要用的数据结构
#   483580 T main.Sort[go.shape.float64_0]
#   483440 T main.Sort[go.shape.float64_0].func1
#   4834a0 T main.Sort[go.shape.int_0]
#   4833e0 T main.Sort[go.shape.int_0].func1
#   483220 T main.main
#   433d80 T runtime.main.func1
#   434100 T runtime.main.func2




#   4ba690 R main..dict.Max[float64]
#   4ba698 R main..dict.Max[int]
#   4ba740 R main..dict.Sort[float64]
#   4ba750 R main..dict.Sort[int]
#   517760 D main..inittask
#   483720 T main.Max[go.shape.float64_0]
#   483700 T main.Max[go.shape.int_0]
#   483620 T main.Sort[go.shape.float64_0]
#   4834e0 T main.Sort[go.shape.float64_0].func1
#   483540 T main.Sort[go.shape.int_0]
#   483480 T main.Sort[go.shape.int_0].func1
#   483220 T main.main
#   433d80 T runtime.main.func1
#   434100 T runtime.main.func2