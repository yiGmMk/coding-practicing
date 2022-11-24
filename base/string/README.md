# strings && bytes

- [strings \&\& bytes](#strings--bytes)
  - [分割](#分割)
    - [Cut](#cut)
      - [参考](#参考)
    - [Split](#split)
    - [SplitN](#splitn)
    - [SplitAfterN](#splitaftern)
  - [拼接](#拼接)
    - [Join](#join)
    - [strings.Builder \&\& bytes.Buffer](#stringsbuilder--bytesbuffer)
  - [修剪/Trim \&\& 变换](#修剪trim--变换)
    - [修剪/Trim](#修剪trim)
      - [TrimSpace](#trimspace)
      - [Trim](#trim)
      - [TrimLeft \& TrimRight](#trimleft--trimright)
      - [TrimPrefix \& TrimSuffix](#trimprefix--trimsuffix)
    - [变换](#变换)
      - [ToUpper/ToLower](#touppertolower)
      - [Map](#map)
    - [对接IO模型](#对接io模型)
      - [strings.NewReader \& bytes.NewReader](#stringsnewreader--bytesnewreader)

## 分割

### Cut

方便分割字符串的实现

```go
// 例如从邮件中取用户名 xxx@xxx.com
idx = strings.Index(username, "@")
if idx != -1 {
  name = username[:idx]
} else {
  name = username
}  
```

```go
// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func Cut(s, sep string) (before, after string, found bool) {
 if i := Index(s, sep); i >= 0 {
  return s[:i], s[i+len(sep):], true
 }
 return s, "", false
}
```

现在可以这么写

```go
user,domain,foud:=Cut(eamil,"@")
```

#### 参考

- [煎鱼 go blog](https://eddycjy.com/posts/go/118-cut/)

### Split

[当传入空串（或bytes.Split被传入nil切片）作为分隔符时，Split函数会按UTF-8的字符编码边界对Unicode进行分割，即每个Unicode字符都会被视为一个分割后的子字符串](./separate_test.go)

```go
fmt.Printf("%q\n", strings.Split("Go社区欢迎你", ""))    // ["G" "o" "社" "区" "欢" "迎" "你"]
fmt.Printf("%q\n", bytes.Split([]byte("Go社区欢迎你"), nil))    // ["G" "o" "社" "区" "欢" "迎" "你"]
```

### SplitN

### SplitAfterN

SplitAfter不同于Split的地方在于它对原字符串/字节切片的分割点在每个分隔符的后面，由于分隔符并未真正起到分隔的作用，因此它不会被剔除，而会作为子串的一部分返回。

```go
fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c,d"), []byte(","), 2))    // ["a," "b,c,d"]
```

## 拼接

### Join

### strings.Builder && bytes.Buffer

```go
s := []string{"I", "love", "Go"}
var builder strings.Builder
for i, w := range s {
    builder.WriteString(w)
    if i != len(s)-1 {
          builder.WriteString(" ")
    }
}
fmt.Printf("%s\n", builder.String()) // I love Go

b := [][]byte{[]byte("I"), []byte("love"), []byte("Go")}
var buf bytes.Buffer
for i, w := range b {
    buf.Write(w)
    if i != len(b)-1 {
        buf.WriteString(" ")
    }
}
fmt.Printf("%s\n", buf.String()) // I love Go
```

## 修剪/Trim && 变换

### 修剪/Trim

#### TrimSpace

TrimSpace函数去除输入字符串/字节切片**首部和尾部**的空白字符
与Fields函数采用的空白字符定义相同,unicode.IsSpace

#### Trim

允许删除自定义字符集合(只修剪首尾字符)

#### TrimLeft & TrimRight

删除首/尾自定义字符集合

#### TrimPrefix & TrimSuffix

删除前/后缀

```go
fmt.Printf("%q\n", strings.TrimLeft("prefix,prefix I love Go!!", "prefix,"))
// " I love Go!!"
fmt.Printf("%q\n", strings.TrimPrefix("prefix,prefix I love Go!!", "prefix,"))
// "prefix I love Go!!"
```

### 变换

#### ToUpper/ToLower

#### Map

```go
trans := func(r rune) rune {
  switch {
  case r == ' ':
   return ','
  }
  return r
 }
 fmt.Printf("%q\n", strings.Map(trans, "CN UK US JP"))
 fmt.Printf("%q\n", bytes.Map(trans, []byte("CN UK US JP")))
 // "CN,UK,US,JP"
 // "CN,UK,US,JP"
```

### 对接IO模型

#### strings.NewReader & bytes.NewReader

[TestIo](./string_test.go)
