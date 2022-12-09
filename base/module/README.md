# go module

- [go module](#go-module)
  - [proxy](#proxy)
    - [proxy设置](#proxy设置)
    - [常用的proxy](#常用的proxy)
    - [package 解析](#package-解析)
    - [自托管Go模块代理](#自托管go模块代理)
    - [参考](#参考)
  - [GOSUMDB](#gosumdb)
  - [GOPRIVATE](#goprivate)
    - [example](#example)
  - [module的主版本号升级](#module的主版本号升级)
    - [修改导入路径](#修改导入路径)
    - [major subdirectory方案](#major-subdirectory方案)

当要使用他提供的package,只需要go get一下就获取到了源代码,这背后是如何实现的?
从哪里下,用什么版本,如何保证安全性,有什么依赖?下面就通过go的Proxy设置一窥.

go module通过GOPROXY,GOPRIVATE,GONOPROXY设置查找package的路径,
用GOSUMDB设置的'检验和数据库'校验package数据的安全,使用最小版本抉择算法:MVS(Minimal version selection)决定需使用的版本.

## proxy

在Go 1.13版本之前，GOPROXY这个环境变量的默认值为空

Go 1.13中将<https://proxy.golang.org设为GOPROXY>环境变量的默认值之一,并开始支持设置多个代理的列表（多个代理之间采用逗号分隔）

Go 1.15版本，新增以管道符“|”为分隔符的代理列表值。如果GOPROXY配置的代理服务列表值以管道符分隔，则无论某个代理服务返回什么错误码，Go命令都会向列表中的下一个代理服务发起新的尝试请求。

### proxy设置

通过环境变量GOPROXY设置proxy

> The go command may be configured to contact proxies or source control servers using the GOPROXY environment variable, which accepts a list of proxy URLs. The list may include the keywords direct or off (see Environment variables for details). List elements may be separated by commas (,) or pipes (|), which determine error fallback behavior. When a URL is followed by a comma, the go command falls back to later sources only after a 404 (Not Found) or 410 (Gone) response. When a URL is followed by a pipe, the go command falls back to later sources after any error, including non-HTTP errors such as timeouts. This error handling behavior lets a proxy act as a gatekeeper for unknown modules. For example, a proxy could respond with error 403 (Forbidden) for modules not on an approved list (see Private proxy serving private modules).

GOPROXY变量

- 支持设置多个proxy
- 默认值是"https://proxy.golang.org,direct"
- 支持以关键字'direct'或'off'标识
  - off: disallows downloading modules from any source
  - off indicates that no communication should be attempted,不请求下载
  - direct: download directly from version control repositories instead of using a module proxy.直接从版本控制系统下载,不使用proxy
- 分隔符决定发生错误时的回退行为,支持的分隔符: ","或"|"
  - 以','分割时,仅当前一个proxy返回错误码404(Not Found)或410(Gone)时才会向下一个proxy发送请求
  - 以'|'分割时,无论发生什么错误,甚至请求超时(非HTTP错误),都会向下一个proxy请求

```sh
export GOPROXY=https://goproxy.cn,direct
```

go 1.15 多proxy,新增以管道符“|”为分隔符的代理列表值,前一个失败后会向下一个proxy发送请求

```sh
export GOPROXY="https://goproxy.cn|https://proxy.golang.com.cn"
```

请求样例(goproxy.cn报错)

```sh
export GOPROXY="https://goproxy.cn|https://proxy.golang.com.cn"

go get -x bitbucket.org/liamstask/goose/cmd/goose
# get https://goproxy.cn/bitbucket.org/@v/list
# get https://goproxy.cn/bitbucket.org/liamstask/goose/cmd/@v/list
# get https://goproxy.cn/bitbucket.org/liamstask/goose/cmd/goose/@v/list
# get https://goproxy.cn/bitbucket.org/liamstask/goose/@v/list
# get https://goproxy.cn/bitbucket.org/liamstask/@v/list
# get https://goproxy.cn/bitbucket.org/liamstask/goose/cmd/goose/@v/list: 500 Internal Server Error (1.960s)
# get https://goproxy.cn/bitbucket.org/liamstask/goose/cmd/@v/list: 500 Internal Server Error (1.972s)
# get https://goproxy.cn/bitbucket.org/@v/list: 500 Internal Server Error (1.990s)
# get https://goproxy.cn/bitbucket.org/liamstask/goose/@v/list: 500 Internal Server Error (2.017s)
# get https://goproxy.cn/bitbucket.org/liamstask/@v/list: 500 Internal Server Error (2.203s)
# get https://proxy.golang.com.cn/bitbucket.org/@v/list
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/cmd/@v/list
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/@v/list
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/@v/list
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/cmd/goose/@v/list
# get https://proxy.golang.com.cn/bitbucket.org/@v/list: 404 Not Found (0.264s)
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/@v/list: 404 Not Found (0.309s)
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/@v/list: 200 OK (0.473s)
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/@latest
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/cmd/goose/@v/list: 404 Not Found (1.185s)
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/cmd/@v/list: 404 Not Found (1.227s)
# get https://proxy.golang.com.cn/bitbucket.org/liamstask/goose/@latest: 200 OK (1.008s)
# get https://goproxy.cn/github.com/ziutek/mymysql/@v/list
# get https://goproxy.cn/github.com/go-sql-driver/mysql/@v/list
# get https://goproxy.cn/github.com/kylelemons/go-gypsy/@v/list
# get https://goproxy.cn/github.com/lib/pq/@v/list
# get https://goproxy.cn/github.com/mattn/go-sqlite3/@v/list
# get https://goproxy.cn/github.com/go-sql-driver/mysql/@v/list: 200 OK (0.116s)
# get https://goproxy.cn/github.com/kylelemons/go-gypsy/@v/list: 200 OK (59.483s)
# get https://goproxy.cn/github.com/lib/pq/@v/list: 200 OK (59.502s)
# get https://goproxy.cn/github.com/ziutek/mymysql/@v/list: 200 OK (59.508s)
# get https://goproxy.cn/github.com/mattn/go-sqlite3/@v/list: 200 OK (59.511s)
go: warning: github.com/mattn/go-sqlite3@v2.0.3+incompatible: retracted by module author: Accidental; no major changes or features.
go: to switch to the latest unretracted version, run:
        go get github.com/mattn/go-sqlite3@latest
```

### 常用的proxy

- proxy.golang.org：Go官方提供的module代理服务。
- mirrors.tencent.com/go：腾讯公司提供的module代理服务。
- mirrors.aliyun.com/goproxy：阿里云提供的module代理服务。
- goproxy.cn：开源module代理，由七牛云提供主机，是目前中国最为稳定的module代理服务。
  支持一系列统计api,如服务摘要,模块趋势,下载次数等,[文档](https://goproxy.cn/stats)
  如:
  - 下载次数
    - [![Goproxy.cn,下载总次数](https://goproxy.cn/stats/github.com/yiGmMk/leetcode/badges/download-count.svg)](https://goproxy.cn)
- goproxy.io：开源module代理，由中国Go社区提供的module代理服务。
- proxy.golang.com.cn:与goproxy.io同一主体提供的代理(大陆地区推荐使用)
- Athens：开源module代理，可基于该代理自行搭建module代理服务。

### package 解析

go加载package时需要明确哪个module内包含那个package.
go先从[build list](https://golang.google.cn/ref/mod#glos-build-list)使用前缀查询包.
查找新的module时会先检查环境变量GOPROXY,根据规则查询:
> A proxy URL indicates the go command should contact a module proxy using the GOPROXY protocol
> direct indicates that the go command should communicate with a version control system.
> off indicates that no communication should be attempted.
> The GOPRIVATE and GONOPROXY environment variables can also be used to control this behavior.

例如,现在GOPROXY设置为"https://corp.example.com,https://proxy.golang.org",需要查找包含golang.org/x/net/html的
module,go的请求行为是这样的

- To <https://corp.example.com/> (in parallel):
  - Request for latest version of golang.org/x/net/html
  - Request for latest version of golang.org/x/net
  - Request for latest version of golang.org/x
  - Request for latest version of golang.org
- To <https://proxy.golang.org/>, if all requests to <https://corp.example.com/> have failed with 404 or 410:
  - Request for latest version of golang.org/x/net/html
  - Request for latest version of golang.org/x/net
  - Request for latest version of golang.org/x
  - Request for latest version of golang.org

查找到一个适用版本后,会在go.mod内新增一个require项(包含module路径,版本)

参考:[Resolving a package to a module](https://golang.google.cn/ref/mod#resolve-pkg-mod)

### 自托管Go模块代理

```go
// 创建一个名为 goproxy.go 的文件
package main

import (
 "net/http"
 "os"

 "github.com/goproxy/goproxy"
)

func main() {
 http.ListenAndServe("localhost:8080", &goproxy.Goproxy{
  GoBinEnv: append(
   os.Environ(),
   "GOPROXY=https://goproxy.cn,direct", // 使用 Goproxy.cn 作为上游代理
   "GOPRIVATE=git.example.com",         // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
  ),
  ProxiedSUMDBs: []string{
   "sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
  },
 })
}

// go run goproxy.go
// 就这么简单，一个功能完备的 Go 模块代理就搭建成功了。
// 事实上，你还可以将 Goproxy 结合着你钟爱的 Web 框架一起使用，比如 Gin 和 Echo，你所需要做的只是多添加一条路由而已。更高级的用法请查看文档。
```

### 参考

- [proxy.io 文档](https://proxy.golang.com.cn/zh/docs/getting-started.html)
- [proxy.cn](https://goproxy.cn/)
- [repo:goproxy](https://github.com/goproxy/goproxy)
- [go doc,module](https://golang.google.cn/ref/mod#private-modules) 或 [go doc,module](https://golang.org/ref/mod)
  - [GOPROXY设置](https://golang.google.cn/ref/mod#environment-variables)
  - [GOPROXY protocol](https://golang.google.cn/ref/mod#module-proxy)
  - [Resolve pkg](https://golang.google.cn/ref/mod#resolve-pkg-mod)

## GOSUMDB

Go 1.13提供了GOSUMDB环境变量来配置Go校验和数据库的服务地址（和公钥），其默认值为"sum.golang.org"

## GOPRIVATE

Go 1.13提供了GOPRIVATE环境变量用于指示哪些仓库下的module是私有的，不需要通过GOPROXY下载，也不需要通过GOSUMDB验证其校验和。

不过要注意的是，GONOPROXY和GONOSUMDB可以覆盖GOPRIVATE变量中的设置

### example

```sh
export GOPRIVATE=github.com/bigwhite/privatemodule
```

## module的主版本号升级

### 修改导入路径

- 在go.mod中升级module的根路径，增加vN
- 建立vN.x.x形式的标签（可选，如果不打标签，Go会在消费者的go.mod中使用伪版本号，比如bitbucket.org/bigwhite/modules-major-branch/v2 v2.0.0-20190603050009-28a5b8da279e）。
- 如果modules-major-branch内部有相互的包引用，那么在升级主版本号的时候，这些包的导入路径也要增加vN，否则就会出现在高版本号的代码中引用低版本号包代码的情况

### major subdirectory方案

利用子目录分割不同主版本,直接用vN作为子目录名字，在代码仓库中将不同版本module放置在不同的子目录中，这样即便不建分支、不打标签，Go编译器通过子目录名也能找到对应的版本
