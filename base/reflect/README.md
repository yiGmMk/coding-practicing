# reflect

reflect包让Go程序具备运行时的反射能力（reflection）。反射是程序在运行时访问、检测和修改它本身状态或行为的一种能力

- [reflect](#reflect)
  - [反射的三大法则](#反射的三大法则)
    - [反射世界的入口](#反射世界的入口)
      - [TypeOf](#typeof)
      - [ValueOf](#valueof)
    - [反射世界的出口](#反射世界的出口)
    - [修改反射对象的前提](#修改反射对象的前提)

## 反射的三大法则

### 反射世界的入口

经由接口（interface{}）类型变量值进入反射的世界并获得对应的反射对象（reflect.Value或reflect.Type）。

#### TypeOf

通过reflect.TypeOf可以得到reflect.Type包含实例的全部类型信息

#### ValueOf

可以得到reflect.Value(包含实例的值信息,还可用过Type方法获取类型信息,[样例](./value_test.go))

### 反射世界的出口

反射对象（reflect.Value）通过化身为一个接口（interface{}）类型变量值的形式走出反射世界。
reflect.Value.Interface()是reflect.ValueOf()的逆过程，通过Interface方法我们可以将reflect.Value对象恢复成一个interface{}类型的变量值

```go
func main() {
    var i = 5
    val := reflect.ValueOf(i)
    r := val.Interface().(int)
    fmt.Println(r) // 5
    r = 6
    fmt.Println(i, r) // 5 6

    val = reflect.ValueOf(&i)
    q := val.Interface().(*int)
    fmt.Printf("%p, %p, %d\n", &i, q, *q) // 0xc0000b4008, 0xc0000b4008, 5
    *q = 7
    fmt.Println(i) // 7
}
```

### 修改反射对象的前提

反射对象对应的reflect.Value必须是可设置的（Settable）
当被反射对象以值类型（T）传递给reflect.ValueOf时，在反射世界中对反射对象值信息的修改不会对被反射对象产生影响。Go的设计者们认为这种修改毫无意义，并禁止了这种行为(panic)

```go
var i = 17
val := reflect.ValueOf(i)
val.SetInt(27) // panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
```

reflect.Value提供了CanSet、CanAddr及CanInterface等方法来帮助我们判断反射对象是否可设置（Settable）、可寻址、可恢复为一个interface{}类型变量
