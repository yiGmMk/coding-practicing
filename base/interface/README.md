# interface

接口类型在运行时的表示分为两部分，一部分是类型信息，一部分是值信息。
只有当接口类型变量的这两部分的值都为nil时，该变量才与nil相等

## 接口的运行时表示

```go
// $GOROOT/src/runtime/runtime2.go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}

type eface struct {
    _type *_type
    data  unsafe.Pointer
}
```

接口类型变量有两种内部表示:eface和iface，这两种表示分别用于不同接口类型的变量。
eface：用于表示没有方法的空接口（empty interface）类型变量，即interface{}类型的变量。
iface：用于表示其余拥有方法的接口（interface）类型变量。

## Trap
