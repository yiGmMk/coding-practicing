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

- [unsafe](#unsafe)
  - [Sizeof](#sizeof)
  - [Alignof](#alignof)
  - [Offsetof](#offsetof)
  - [Add](#add)
  - [Pointer](#pointer)
    - [安全使用Pointer](#安全使用pointer)
      - [Conversion of a *T1 to Pointer to*T2](#conversion-of-a-t1-to-pointer-tot2)
      - [Conversion of a Pointer to a uintptr (but not back to Pointer)](#conversion-of-a-pointer-to-a-uintptr-but-not-back-to-pointer)
      - [Conversion of a Pointer to a uintptr and back, with arithmetic](#conversion-of-a-pointer-to-a-uintptr-and-back-with-arithmetic)
      - [Conversion of a Pointer to a uintptr when calling syscall.Syscall](#conversion-of-a-pointer-to-a-uintptr-when-calling-syscallsyscall)
      - [Conversion of the result of reflect.Value.Pointer or reflect.Value.UnsafeAddr from uintptr to Pointer](#conversion-of-the-result-of-reflectvaluepointer-or-reflectvalueunsafeaddr-from-uintptr-to-pointer)
      - [Conversion of a reflect.SliceHeader or reflect.StringHeader Data field to or from Pointer](#conversion-of-a-reflectsliceheader-or-reflectstringheader-data-field-to-or-from-pointer)

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

参考go unsafe源码中的几种模式使用.

#### Conversion of a *T1 to Pointer to*T2

允许转换成不同类型数据,需要注意的是T1,T2的memory layout应该是相当的.

转换后类型T2的对齐系数不能比转换前类型T1的对齐系数更严格，即Alignof(T1) >= Alignof(T2)

```go
// Provided that T2 is no larger than T1 and that the two share an equivalent(等价/相等的)
// memory layout, this conversion allows reinterpreting data of one type as
// data of another type. An example is the implementation of
// math.Float64bits:

func Float64bits(f float64) uint64 {
 return *(*uint64)(unsafe.Pointer(&f))
}
```

#### Conversion of a Pointer to a uintptr (but not back to Pointer)

A uintptr is an integer, not a **reference**.

```go
// Converting a Pointer to a uintptr produces the memory address of the value
// pointed at, as an integer. The usual use for such a uintptr is to print it.
//
// Conversion of a uintptr back to Pointer is not valid in general.
//
// A uintptr is an integer, not a reference.
// Converting a Pointer to a uintptr creates an integer value
// with no pointer semantics.
// Even if a uintptr holds the address of some object,
// the garbage collector will not update that uintptr's value
// if the object moves, nor will that uintptr keep the object
// from being reclaimed.
//
// The remaining patterns enumerate the only valid conversions
// from uintptr to Pointer.
//
```

#### Conversion of a Pointer to a uintptr and back, with arithmetic

arithmetic,算术运算
常用于获取struct中的field/字段,注意不要越界

```go
// If p points into an allocated object, it can be advanced through the object
// by conversion to uintptr, addition of an offset, and conversion back to Pointer.
//
p = unsafe.Pointer(uintptr(p) + offset)
//
// The most common use of this pattern is to access fields in a struct
// or elements of an array:
//
// equivalent to f := unsafe.Pointer(&s.f)
f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))
//
// // equivalent to e := unsafe.Pointer(&x[i])
e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))
//
// It is valid both to add and to subtract offsets from a pointer in this way.
// It is also valid to use &^ to round pointers, usually for alignment.
// In all cases, the result must continue to point into the original allocated object.
//
// Unlike in C, it is not valid to advance a pointer just beyond the end of
// its original allocation:
//
// // INVALID: end points outside allocated space.
var s thing
end = unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))
//
// // INVALID: end points outside allocated space.
b := make([]byte, n)
end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))

// Note that both conversions must appear in the same expression, with only
// the intervening arithmetic between them:
//
// INVALID: uintptr cannot be stored in variable
// before conversion back to Pointer.
u := uintptr(p)
p = unsafe.Pointer(u + offset)

// Note that the pointer must point into an allocated object, so it may not be nil.
//
// INVALID: conversion of nil pointer
u := unsafe.Pointer(nil)
p := unsafe.Pointer(uintptr(u) + offset)
```

#### Conversion of a Pointer to a uintptr when calling syscall.Syscall

使用系统调用(syscall)时将Pointer转成uintptr

```go
// The Syscall functions in package syscall pass their uintptr arguments directly
// to the operating system, which then may, depending on the details of the call,
// reinterpret some of them as pointers.
// That is, the system call implementation is implicitly converting certain arguments
// back from uintptr to pointer.
//
// If a pointer argument must be converted to uintptr for use as an argument,
// that conversion must appear in the call expression itself:
//
syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))
//
// The compiler handles a Pointer converted to a uintptr in the argument list of
// a call to a function implemented in assembly by arranging that the referenced
// allocated object, if any, is retained and not moved until the call completes,
// even though from the types alone it would appear that the object is no longer
// needed during the call.
//
// For the compiler to recognize this pattern,
// the conversion must appear in the argument list:
//
// // INVALID: uintptr cannot be stored in variable
// // before implicit conversion back to Pointer during system call.
u := uintptr(unsafe.Pointer(p))
syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))
```

#### Conversion of the result of reflect.Value.Pointer or reflect.Value.UnsafeAddr from uintptr to Pointer

注意在一个表达式中完成转换,不要将返回值赋值给一个uintptr类型变量再在后续的语句中进行转换。

```go
// Package reflect's Value methods named Pointer and UnsafeAddr return type uintptr
// instead of unsafe.Pointer to keep callers from changing the result to an arbitrary
// type without first importing "unsafe". However, this means that the result is
// fragile and must be converted to Pointer immediately after making the call,
// in the same expression:
//
p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
//
// As in the cases above, it is invalid to store the result before the conversion:
//
// INVALID: uintptr cannot be stored in variable
// before conversion back to Pointer.
u := reflect.ValueOf(new(int)).Pointer()
p := (*int)(unsafe.Pointer(u))
```

#### Conversion of a reflect.SliceHeader or reflect.StringHeader Data field to or from Pointer

```go
// As in the previous case, the reflect data structures SliceHeader and StringHeader
// declare the field Data as a uintptr to keep callers from changing the result to
// an arbitrary type without first importing "unsafe". However, this means that
// SliceHeader and StringHeader are only valid when interpreting the content
// of an actual slice or string value.
//
var s string
hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
hdr.Len = n
//
// In this usage hdr.Data is really an alternate(交替的/轮流的/间隔的;代理人/代替者) way to refer to the underlying
// pointer in the string header, not a uintptr variable itself.
//
// In general, reflect.SliceHeader and reflect.StringHeader should be used
// only as *reflect.SliceHeader and*reflect.StringHeader pointing at actual
// slices or strings, never as plain structs.
// A program should not declare or allocate variables of these struct types.
// INVALID: a directly-declared header will not hold Data as a reference.
var hdr reflect.StringHeader
hdr.Data = uintptr(unsafe.Pointer(p))
hdr.Len = n
s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lost
```
