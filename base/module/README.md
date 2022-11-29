# go module

## proxy

在Go 1.13版本之前，GOPROXY这个环境变量的默认值为空

Go 1.13中将<https://proxy.golang.org设为GOPROXY>环境变量的默认值之一,并开始支持设置多个代理的列表（多个代理之间采用逗号分隔）

Go 1.15版本，新增以管道符“|”为分隔符的代理列表值。如果GOPROXY配置的代理服务列表值以管道符分隔，则无论某个代理服务返回什么错误码，Go命令都会向列表中的下一个代理服务发起新的尝试请求。

## GOSUMDB

Go 1.13提供了GOSUMDB环境变量来配置Go校验和数据库的服务地址（和公钥），其默认值为"sum.golang.org"

## GOPRIVATE

Go 1.13提供了GOPRIVATE环境变量用于指示哪些仓库下的module是私有的，不需要通过GOPROXY下载，也不需要通过GOSUMDB验证其校验和。

不过要注意的是，GONOPROXY和GONOSUMDB可以覆盖GOPRIVATE变量中的设置

## module的主版本号升级

### 修改导入路径

- 在go.mod中升级module的根路径，增加vN
- 建立vN.x.x形式的标签（可选，如果不打标签，Go会在消费者的go.mod中使用伪版本号，比如bitbucket.org/bigwhite/modules-major-branch/v2 v2.0.0-20190603050009-28a5b8da279e）。
- 如果modules-major-branch内部有相互的包引用，那么在升级主版本号的时候，这些包的导入路径也要增加vN，否则就会出现在高版本号的代码中引用低版本号包代码的情况

### major subdirectory方案

利用子目录分割不同主版本,直接用vN作为子目录名字，在代码仓库中将不同版本module放置在不同的子目录中，这样即便不建分支、不打标签，Go编译器通过子目录名也能找到对应的版本
