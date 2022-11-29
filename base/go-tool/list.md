# go list

用于列出关于包/module的各类信息

在module-aware模式,会在当前路径下寻找go.mod

- [go list](#go-list)
  - [列出当前路径及其子路径（递归）下的所有包](#列出当前路径及其子路径递归下的所有包)
  - [go保留的路径关键字](#go保留的路径关键字)
    - [main](#main)
    - [all](#all)
    - [cmd](#cmd)
    - [std](#std)
  - [-m](#-m)
  - [格式化 -f](#格式化--f)
  - [-json 以json格式输出](#-json-以json格式输出)
  - [module的可用升级版本](#module的可用升级版本)

## 列出当前路径及其子路径（递归）下的所有包

可以用go list {当前路径}/...

```sh
# module-aware模式
GO111MODULE=on go list ./...     
#gopath模式
GO111MODULE=off go list ./...     
```

也可以使用导入路径加...列出该路径下所有包导入路径

```sh
GO111MODULE=on go list github.com/yiGmMk/leetcode/... 
GO111MODULE=off go list github.com/yiGmMk/leetcode/... 
```

## go保留的路径关键字

### main

表示独立可执行程序的顶层包

### all

- 在gopath模式下，它可以展开为标准库和GOPATH路径下的所有包
- 在module-aware模式下，它展开为主module（当前路径下的module）下的所有包及其所有依赖包，包括测试代码的依赖包

### cmd

代码Go语言自身项目仓库下的src/cmd下的所有包及internal包

### std

代表标准库所有包的集合

## -m

列出module信息

## 格式化 -f

-f格式字符串采用go template包的语法,go list等价于

```go
go list -f '{{.ImportPath}}'
```

支持的字段来自runtime下的PackagePublic

```go
// $GOROOT/src/cmd/go/internal/pkg.go (go 1.14)
type PackagePublic struct {
    Dir           string `json:",omitempty"`  // 包含包源码的目录
    ImportPath    string `json:",omitempty"`  // dir下包的导入路径
    ImportComment string `json:",omitempty"`  // 包声明语句后面的注释中的路径
    Name          string `json:",omitempty"`  // 包名
    Doc           string `json:",omitempty"`  // 包文档字符串
    Target        string `json:",omitempty"`  // 该软件包的安装目标（可以是可执行的）
    ...

    TestGoFiles  []string `json:",omitempty"` // 包中的_test.go文件
    TestImports  []string `json:",omitempty"` // TestGoFiles导入的包
    XTestGoFiles []string `json:",omitempty"` // 包外的_test.go
    XTestImports []string `json:",omitempty"` // XTestGoFiles导入的包
}
```

```sh
# 输出包相关信息
go list -f '{{printf "%#v" .}}'
```

## -json 以json格式输出

```sh
go list -json
```

## module的可用升级版本

```sh
go list -m -u all
```
