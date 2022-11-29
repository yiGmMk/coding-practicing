# go lint tool

- [go lint tool](#go-lint-tool)
  - [gofmt](#gofmt)
  - [goimports](#goimports)
  - [go vet](#go-vet)
    - [assign规则](#assign规则)
    - [atomic规则](#atomic规则)
    - [buildtag规则](#buildtag规则)
    - [composites规则](#composites规则)
    - [copylocks规则](#copylocks规则)
    - [loopclosure规则](#loopclosure规则)
    - [unsafeptr规则](#unsafeptr规则)
    - [unmarshal规则](#unmarshal规则)
  - [golangci-lint](#golangci-lint)

## gofmt

## goimports

## go vet

内置多条静态代码检查规则.

```sh
#支持手动关闭规则
go vet -printf=false -buildtag=false xx.go
```

### assign规则

无效赋值操作

### atomic规则

检查sync.Atomic包函数误用

### buildtag规则

检查源文件中+build tag是否正确定义

### composites规则

检查源文件中是否有未使用“field:value”格式的复合字面值形式对struct类型变量进行值构造的问题

### copylocks规则

检查源文件中是否存在lock类型变量的按值传递问题

### loopclosure规则

检查源文件中是否存在循环内的匿名函数引用循环变量的问题

### unsafeptr规则

检查源码中是否有非法将uintptr转换为unsafe.Pointer的问题

### unmarshal规则

检查源码中是否有将非指针或非接口类型值传给unmarshal的问题

## golangci-lint

```sh
# 获取所有支持的linter
golangci-lint linters
```
