# refactoring/重构

## gofmt -r 纯字符串替换

默认写到标准输出，这种类似干跑（dry run）的方式可以用于评估替换效果

```sh
gofmt -r '"log" -> "github.com/rs/zerolog/log"' -w .
```

加上-w 写入文件

## gorename 安全的标识符替换(重命名)

## gomvpkg 移动包并更新包导入路径
