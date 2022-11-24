# unsafe

```go
// go 1.19 下unsafe包
type ArbitraryType int
type IntegerType int
type Pointer *ArbitraryType

func Sizeof(x ArbitraryType) uintptr   
func Offsetof(x ArbitraryType) uintptr 
func Alignof(x ArbitraryType) uintptr  
func Add(ptr Pointer, len IntegerType) Pointer
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
```

## Sizeof

获取一个表达式值的大小

## Alignof

用于获取一个表达式的内存地址对齐系数
对齐系数（alignment factor）是一个计算机体系架构（computer architecture）层面的术语。
在不同的计算机体系结构下，处理器对变量地址都有着对齐要求，即变量的地址必须可被该变量的对齐系数整除

## Offsetof

获取结构体中某字段的地址偏移量（相对于结构体变量的地址）

## Add

## Pointer

```go
// Pointer represents a pointer to an arbitrary type. There are four special operations
// available for type Pointer that are not available for other types:
//   - A pointer value of any type can be converted to a Pointer.
//     任意类型的指针值都可以被转换为unsafe.Pointer

//   - A Pointer can be converted to a pointer value of any type.
//     也可以被转换为任意类型的指针值

//   - A uintptr can be converted to a Pointer.
//     uintptr类型值可以被转换为一个unsafe.Pointer

//   - A Pointer can be converted to a uintptr.
//     unsafe.Pointer也可以被转换为一个uintptr类型值
```

由于以上4个特性,使用Pointer能够突破go的强类型保护,任意修改内存

### 安全使用Pointer

