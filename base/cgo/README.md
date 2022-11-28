# cgo

```go
package main

// #include <stdio.h>
// #include <stdlib.h>
//
// void print(char *str) {
//     printf("%s\n", str);
// }
import "C"

import (
 "unsafe"
)

func main() {
 s := "Hello, Cgo"
 cs := C.CString(s)
 defer C.free(unsafe.Pointer(cs))
 C.print(cs) // Hello, Cgo
}

```

Go源码文件中的C代码是需要用注释包裹的，就像上面的include头文件以及print函数定义那样。其次，import "C"这行语句是必须有的，而且其与上面的C代码之间不能用空行分隔，必须紧密相连。这里的“C”不是包名，而是一种类似名字空间的概念，也可以理解为伪包名，C语言所有语法元素均在该伪包下面。

## 静态编译

如果使用了net这样的包含cgo实现版本的标准库包，那么CGO_ENABLED的值将影响你的程序编译后的属性（是静态的还是动态链接的）
如果CGO_ENABLED=1且仅使用了net、os/user等依赖cgo实现的包，那么internal linking机制将被默认采用，编译过程不会采用静态链接；但如若依然要强制静态编译，需向-ldflags '-linkmode "external" -extldflags "-static"'传递go build命令

```sh
$go build -o how_cgo_works_static  -ldflags '-extldflags "-static"' how_cgo_works.go
$ldd how_cgo_works_static
    not a dynamic executable
```

默认情况下，Go的运行时环境变量CGO_ENABLED=1（通过go env命令可以查看），即默认开启cgo，允许你在Go代码中调用C代码.默认情况下不使用静态编译,会有动态库依赖.

```go
func main() {
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    srv := &http.Server{
        Addr:    ":8000",
        Handler: http.FileServer(http.Dir(cwd)),
    }
    log.Fatal(srv.ListenAndServe())
}
```

```sh
$go build -o server_default server.go
$ldd server_default
    linux-vdso.so.1 (0x00007ffde17ef000)
    libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f721b0f4000)
    libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f721ad03000)
    /lib64/ld-linux-x86-64.so.2 (0x00007f721b313000)
```
