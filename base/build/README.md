# build

## cmd param

### -x -v

go build并不会输出执行了哪些命令以及这些命令的执行细节，我们能看到的要么是构建成功，要么是go build输出的一些
错误信息
通过传入-v和-x命令行标志选项让go build输出构建过程构建了哪些包，执行了哪些命令.

```sh
WORK=/tmp/go-build3767719814
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=/home/admin/.cache/go-build/dc/dc33c177be764b0e77d936f42e3599c52a204647fefca2bca728784e99a7068d-d
packagefile fmt=/usr/local/go/pkg/linux_amd64/fmt.a
packagefile github.com/duke-git/lancet/v2/lancetconstraints=/home/admin/.cache/go-build/97/97fdecfb831e4a4d0c011f3cac034b9c9a9367c7a603556879a2ce173b82c4f9-d
packagefile sort=/usr/local/go/pkg/linux_amd64/sort.a
packagefile runtime=/usr/local/go/pkg/linux_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/linux_amd64/errors.a
packagefile internal/fmtsort=/usr/local/go/pkg/linux_amd64/internal/fmtsort.a
packagefile io=/usr/local/go/pkg/linux_amd64/io.a
packagefile math=/usr/local/go/pkg/linux_amd64/math.a
packagefile os=/usr/local/go/pkg/linux_amd64/os.a
packagefile reflect=/usr/local/go/pkg/linux_amd64/reflect.a
packagefile strconv=/usr/local/go/pkg/linux_amd64/strconv.a
packagefile sync=/usr/local/go/pkg/linux_amd64/sync.a
packagefile unicode/utf8=/usr/local/go/pkg/linux_amd64/unicode/utf8.a
packagefile internal/reflectlite=/usr/local/go/pkg/linux_amd64/internal/reflectlite.a
packagefile math/bits=/usr/local/go/pkg/linux_amd64/math/bits.a
packagefile internal/abi=/usr/local/go/pkg/linux_amd64/internal/abi.a
packagefile internal/bytealg=/usr/local/go/pkg/linux_amd64/internal/bytealg.a
packagefile internal/cpu=/usr/local/go/pkg/linux_amd64/internal/cpu.a
packagefile internal/goarch=/usr/local/go/pkg/linux_amd64/internal/goarch.a
packagefile internal/goexperiment=/usr/local/go/pkg/linux_amd64/internal/goexperiment.a
packagefile internal/goos=/usr/local/go/pkg/linux_amd64/internal/goos.a
packagefile runtime/internal/atomic=/usr/local/go/pkg/linux_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/usr/local/go/pkg/linux_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/usr/local/go/pkg/linux_amd64/runtime/internal/sys.a
packagefile runtime/internal/syscall=/usr/local/go/pkg/linux_amd64/runtime/internal/syscall.a
packagefile internal/itoa=/usr/local/go/pkg/linux_amd64/internal/itoa.a
packagefile internal/oserror=/usr/local/go/pkg/linux_amd64/internal/oserror.a
packagefile internal/poll=/usr/local/go/pkg/linux_amd64/internal/poll.a
packagefile internal/syscall/execenv=/usr/local/go/pkg/linux_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/usr/local/go/pkg/linux_amd64/internal/syscall/unix.a
packagefile internal/testlog=/usr/local/go/pkg/linux_amd64/internal/testlog.a
packagefile internal/unsafeheader=/usr/local/go/pkg/linux_amd64/internal/unsafeheader.a
packagefile io/fs=/usr/local/go/pkg/linux_amd64/io/fs.a
packagefile sync/atomic=/usr/local/go/pkg/linux_amd64/sync/atomic.a
packagefile syscall=/usr/local/go/pkg/linux_amd64/syscall.a
packagefile time=/usr/local/go/pkg/linux_amd64/time.a
packagefile unicode=/usr/local/go/pkg/linux_amd64/unicode.a
packagefile internal/race=/usr/local/go/pkg/linux_amd64/internal/race.a
packagefile path=/usr/local/go/pkg/linux_amd64/path.a
modinfo "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\ndep\tgithub.com/duke-git/lancet/v2\tv2.0.7\th1:vDZH3NCtTGD8vVytxRmmgHfAOEhNLua2qrxqMY2/8Ew=\nbuild\t-compiler=gc\nbuild\tCGO_ENABLED=1\nbuild\tCGO_CFLAGS=\nbuild\tCGO_CPPFLAGS=\nbuild\tCGO_CXXFLAGS=\nbuild\tCGO_LDFLAGS=\nbuild\tGOARCH=amd64\nbuild\tGOOS=linux\nbuild\tGOAMD64=v1\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=O_RNGTWnWjU3dVzRfYfC/GPKJ2S6VBRd6ikpCkvry/xq28n1EVWqOtf-EZpPwy/O_RNGTWnWjU3dVzRfYfC -extld=gcc /home/admin/.cache/go-build/dc/dc33c177be764b0e77d936f42e3599c52a204647fefca2bca728784e99a7068d-d
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out main
rm -r $WORK/b001/
```

- 创建用于构建的临时目录；
- 下载构建依赖的module；
- 分别编译依赖的module将编译后的结果存储到临时目录及GOCACHE目录下；
- 编译目标module；
- 定位和汇总目标module的各个依赖包构建后的目标文件（.a文件）的位置，形成importcfg.link文件，供后续链接器使用；
- 链接成可执行文件；
- 清理临时构建环境。

### -o

指定输出文件位置,名称

```sh
go build -o main xxx.go
```

### -race

在构建的结果中加入竞态检测的代码。在程序运行过程中，如果发现对数据的并发竞态访问，这些代码会给出警告，这些警告信息可以用来辅助后续查找和解决竞态问题

### -gcflags 传给编译器的标志

```sh
go build -gcflags[=标志应用的包范围]='空格分隔的标志选项列表'
```

其中“标志应用的包范围”是可选项，如果不显式填写，那么Go编译器仅将通过gcflags传递的编译选项应用在当前包；如果显式指定了包范围，则通过gcflags传递的编译选项不仅会应用在当前包的编译上，还会应用于包范围指定的包上

```go
go build -gcflags='-N -l'      // 仅将传递的编译选项应用于当前包
go build -gcflags=all='-N -l'  // 将传递的编译选项应用于当前包及其所有依赖包
go build -gcflags=std='-N -l'  // 仅将传递的编译选项应用于标准库包

go build -gcflags='-m'
go build -gcflags='-m -m'      // 输出比上一条命令更为详尽的逃逸分析过程信息
go build -gcflags='-m=2'       // 与上一条命令等价
go build -gcflags='-m -m -m'   // 输出最为详尽的逃逸分析过程信息
go build -gcflags='-m=3'       // 与上一条命令等价
```

- -l：关闭内联。
- -N：关闭代码优化。
- -m：输出逃逸分析（决定哪些对象在栈上分配，哪些对象在堆上分配）的分析决策过程。
- -S：输出汇编代码。

### -ldflags

```sh
go tool link -help
```

- -X 设定包中string类型变量的值（仅支持string类型变量）
 通过-X选项，我们可以在编译链接期间动态地为程序中的字符串变量进行赋值，这个选项的一个典型应用就是在构建脚本中设定程序的版本值。我们通常会为应用程序添加version命令行标志选项，用来输出当前程序的版本信息，就像Go自身那样：

 ```go
 $go version
 go version go1.14 darwin/amd64
 
 
 var (  
       version string
     )
 func main() { 
     if os.Args[1] == "version" {  
         fmt.Println("version:", version)       
         return  
     }
 }
 ```

 在这个源文件中，我们并未显式初始化version这个变量。接下来，构建这个程序，在构建时为version这个string类型变量动态地注入新值

 ```sh
 $go build -ldflags "-X main.version=v0.7.0" linker_x_flag.go
 $./linker_x_flag version
 version: v0.7.0
 ```

- -s：不生成符号表（symbol table）
- -w：不生成DWARF（Debugging With Attributed Record Formats）调试信息

### -tags：指定构建约束条件

go build可以通过-tags指定构建的约束条件，以决定哪些源文件被包含在包内进行构建。tags的值为一组逗号分隔（老版本为空格分隔）的值

```go
$go build -tags="tag1,tag2,..." ...
```

build tag行也是注释，它以+build作为起始标记，与前面的注释符号//中间有一个空格。+build后面就是约束标记字符串,每一行的build tag实质上会被求值为一个布尔表达式.

```go
// 表达式                           含义
// +build t1,t2                    t1 and t2
// +build t1 t2                    t1  or t2
// +build !t2                      not t2
// 多行之间是与的关系, and
```
