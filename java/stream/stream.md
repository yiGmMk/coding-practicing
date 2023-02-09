# stream

Stream提供的常用操作有

- 转换操作：map()，filter()，sorted()，distinct()
- 合并操作：concat()，flatMap()
- 并行处理：parallel()
- 聚合操作：reduce()，collect()，count()，max()，min()，sum()，average()
- 其他操作：allMatch(), anyMatch(), forEach()

## map

## filter

## reduce

## 其他

### 排序

要求Stream的每个元素必须实现Comparable接口。如果要自定义排序，传入指定的Comparator即可

```java

Stream.of("b", "a", "d", "c").sorted()
    .collect(Collectors.toList()).forEach(System.out::println);
```

### 截取

截取操作常用于把一个无限的Stream转换成有限的Stream，
skip()用于跳过当前Stream的前N个元素，limit()用于截取当前Stream最多前N个元素

```java
Stream.of("B", "a", "D", "c")
                .sorted(String::compareToIgnoreCase)
                .limit(2)
                .skip(1)
                .collect(Collectors.toList()).forEach(System.out::println);
// B
```

## 参考

- <https://www.liaoxuefeng.com/wiki/1252599548343744/1322403061825570>
