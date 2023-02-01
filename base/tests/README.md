# go 测试之道

- [go 测试之道](#go-测试之道)
  - [包内测试/包外测试](#包内测试包外测试)
    - [包外测试](#包外测试)
      - [后门](#后门)
  - [测试套件,xUnit](#测试套件xunit)
  - [表驱动测试(table driver test)](#表驱动测试table-driver-test)
  - [testdata,测试依赖的外部数据管理](#testdata测试依赖的外部数据管理)
    - [golden文件](#golden文件)
  - [fake,stub,mock](#fakestubmock)
    - [fake,真实组件或服务的简化实现替身](#fake真实组件或服务的简化实现替身)
    - [stub,对返回结果有一定预设控制能力的替身](#stub对返回结果有一定预设控制能力的替身)
    - [mock,专用于行为观察和验证的替身](#mock专用于行为观察和验证的替身)
      - [go官方维护的mock工具](#go官方维护的mock工具)
  - [fuzz](#fuzz)
  - [benckmark,性能测试](#benckmark性能测试)

《Go语言精进之路:从新手到高手》在"测试"这一主题相关的部分章节,
作者从测试文件的位置(包内/包外测试)讲起,介绍了测试代码的层次结构,常见测试模式,表驱动测试和
go对测试数据组织做的优化,以及mock测试等辅助单元测试,fuzz测试

## 包内测试/包外测试

包内测试:测试代码与被测代码在同一包内

```bash
#包外测试
go list -f={{.XTestGoFiles}} .
#在$GoROOT/src/strings目录下为 [builder_test.go clone_test.go compare_test.go example_test.go reader_test.go replace_test.go search_test.go strings_test.go]

#包内测试,在$GoROOT/src/strings目录下为 [export_test.go]
go list -f={{.TestGoFiles}} .
```

### 包外测试

仅针对导出API进行测试,是面向接口的黑盒测试.

- 与被测代码充分解耦,方便维护
- 解决"包循环引用"问题
- 从用户视角验证导出API的合理性和易用性
- 对被测代码覆盖不足

#### 后门

针对包外测试对被测代码覆盖不足的情况,可以使用"安插后门"的方式,即在测试文件(*_test.go)提供代码用于测试

在测试文件中的代码不对包含在正式产品代码中.

## 测试套件,xUnit

利用go在1.14后提供的testing.CleanUp,实现测试固件的销毁,TestMain实现包级别的测试固件创建和销毁.

参考testsuite下的测试案例

在 go中也能实现类似Java,Pyhton等其他语言里主流的xUnit类似的测试套件(Test Suite)

```go
package strings_test
func testBuilder(t *testing.T) {
    ...
}
func testBuilderString(t *testing.T) {
    ...
}
func testBuilderReset(t *testing.T) {
    ...
}
func testBuilderGrow(t *testing.T) {
    ...
}
func TestBuilder(t *testing.T) {
    t.Run("TestBuilder", testBuilder)
    t.Run("TestBuilderString", testBuilderString)
    t.Run("TestBuilderReset", testBuilderReset)
    t.Run("TestBuilderGrow", testBuilderGrow)
}
```

## 表驱动测试(table driver test)

基于数据的测试

```go
func TestCompare(t *testing.T) {
    compareTests := []struct {
        name, a, b string
        i          int
    }{
        {`compareTwoEmptyString`, "", "", 0},
        {`compareSecondParamIsEmpty`, "a", "", 1},
        {`compareFirstParamIsEmpty`, "", "a", -1},
    }

    for _, tt := range compareTests {
        t.Run(tt.name, func(t *testing.T) {
            cmp := strings.Compare(tt.a, tt.b)
            if cmp != tt.i {
                t.Errorf(`want %v, but Compare(%q, %q) = %v`, tt.i, tt.a, tt.b, cmp)
            }
        })
    }
}
```

还可以单独执行某项数据的测试

```sh
$go test -v -run /TwoEmptyString ./base/tests/table_driven_test.go
=== RUN   TestCompare
=== RUN   TestCompare/compareTwoEmptyString
--- PASS: TestCompare (0.00s)
    --- PASS: TestCompare/compareTwoEmptyString (0.00s)
PASS
ok     command-line-arguments   0.005s
```

## testdata,测试依赖的外部数据管理

go工具链默认忽略testdata目录.

### golden文件

将预期数据采集到文件的过程与测试代码结合

## fake,stub,mock

fake,stub,mock概念,区分可参考[xUnit Test Patterns](https://book.douban.com/subject/1859393)

### fake,真实组件或服务的简化实现替身

使用真实主键或简化版本作为被测代码的外部依赖,如测试里常用的redis包
github.com/alicebob/miniredis/v2,内存中运行的redis

### stub,对返回结果有一定预设控制能力的替身

### mock,专用于行为观察和验证的替身

mock替身更为强大,除了能提供测试前的预设置返回结果能力之外,还可以对mock替身对象在测试过程中的行为进行观察和验证

#### go官方维护的mock工具

[gomock](https://github.com/golang/mock),支持用mockgen生成代码(用于mock interface,即接口)
mockgen支持source/reflect两种方式生成mock代码

## fuzz

## benckmark,性能测试
