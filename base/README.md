# go

## 版本切换

### goup

#### 安装

```shell
curl -sSf https://raw.githubusercontent.com/owenthereal/goup/master/install.sh | sh -s -- '--skip-prompt'
```

#### 使用

安装在~/.go/{version}/下,如go1.13.9在~/.go/go1.13.9/bin/go

### go 官方方案

```shell
go get golang.org/dl/go{version}

go{version} download
```

参考:

[go get golang.org/dl/go1.16.4](https://polarisxu.studygolang.com/posts/go/managing-multiple-go-versions/)