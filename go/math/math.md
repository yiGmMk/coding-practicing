# 数学

## 位运算

### n & (n-1)

#### 判断一个数是否是2的次幂

```go
check := func(x int) bool {
  if x > 0 && x&(x-1) == 0 {
   return true
  }
  return false
}
```

## 排列组合
