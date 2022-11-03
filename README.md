# coding & practicing

## design patterns

![design patterns](./design-pattern/view.PNG)

### references

- [refactoringguru.cn](https://refactoringguru.cn/design-patterns/abstract-factory)
- [博文,设计模式](https://lailin.xyz/post/factory.html)
- [博文,刘丹冰aceld,设计模式](https://www.yuque.com/aceld/lfhu8y/pebesh)

## algorithm & datastructure

### algorithm

sort/search/bitmap

### datastructure

tree/list/array

## awesome xxx & mark

### collection

1. [awesome-go-cn,Go框架、库和软件的中文收录大全](https://github.com/yinggaozhen/awesome-go-cn)
2. [开发者头条,博客/公众号各种站点资讯,提供订阅服务](https://toutiao.io/posts/hot/30)
3. [架构师](https://github.com/xingshaocheng/architect-awesome)
4. [github 资讯](https://github.com/GitHubDaily/GitHubDaily)
5. [repo: algorithm & data structure](https://github.com/KeKe-Li/data-structures-questions)
6. [repo: it资料大全](https://github.com/0voice)

### database

- [mysql 8.0 document](https://dev.mysql.com/doc/refman/8.0/en/preface.html)
- [repo: TiDB Talent-Plan,分布式数据库,分布式系统](https://github.com/pingcap/talent-plan)
- [极客时间 丁奇 Mysql45讲](https://time.geekbang.org/column/intro/100020801)
- [SQL 审核查询平台](https://github.com/hhyo/Archery)
- [mysql 日期相关函数](https://blog.csdn.net/hu1010037197/article/details/115391335)
- [知乎专栏,mysql 优化](https://www.zhihu.com/column/c_1439702957097447424)
- [TiDB计数月刊](https://tidb.net/book/tidb-monthly/2022-09/)
- [repo: 电子科技大学,分布式存储](<https://github.com/CDDSCLab/training-plan>)

#### 博文

- [LeetCode题解 mysql 查询策略,窗口函数,排序](https://leetcode.cn/problems/nth-highest-salary/solution/mysql-zi-ding-yi-bian-liang-by-luanz/)
- [mysql 窗口函数(Window Function),文档](https://dev.mysql.com/doc/refman/8.0/en/window-function-descriptions.html)
- [MySQL TIME_WAIT连接,2MSL](https://blog.csdn.net/fkew2009/article/details/86714699)

### blog

1. [小林coding](https://xiaolincoding.com/)

### notes

- [笔记,汇总 programnotes.cn](https://programnotes.cn/a-set-of-notes-blogs/index.html)
- 正则表达式

### linux

1. [shell配置,bashrc,profile](https://zhuanlan.zhihu.com/p/405174594)

### language

#### python

- [doc](https://docs.python.org/zh-cn/3.10/tutorial/index.html)

#### go

- [effective go](https://golang.google.cn/doc/effective_go.html)
- [go code review](https://github.com/golang/go/wiki/CodeReviewComments#gofmt)
- [go basic,Golang基础框架图](image/go/Golang基础框架图.png)

##### profile

profile通常分为:追踪型(追踪提前设定的事件,如函数调用,含税退出)/采样型

go支持的profile有:cpu/memory/block/goroutine

go的cpu profile在Linux系统使用信号中断(SIGPROF signal)采集运行数据.
>SIGPROF signal: This signal typically indicates expiration of a timer that measures both CPU time used by the current process, and CPU time expended on behalf of the process by the system. Such a timer is used to implement code profiling facilities, hence the name of this signal.
go通过SIGPROF注册回调函数每10ms采集一次.
**需要注意的是:统计时间与用户体验到的时间通常不同**,[profile notes](https://github.com/DataDog/go-profiler-notes/blob/main/guide/README.md),例如一次http请求耗时100ms(数据库耗时95ms,cpu5ms)

memory profile同样基于采样生成

- [go doc,profile&trace](h)
- [profiling go programs](https://golang.google.cn/blog/pprof)
- [profile](https://mp.weixin.qq.com/s/DRQWcU2dN-FycoyFZfnklA) go profile原理
- [Go Profiler Internals](https://www.instana.com/blog/go-profiler-internals/)
- [repo go-profiler-notes](https://github.com/DataDog/go-profiler-notes) go profile

#### repo & tool

- [apisix](https://apisix.apache.org/zh/blog/2021/05/24/tencent-games/)
- [Introduction-to-Golang](https://github.com/0voice/Introduction-to-Golang)

## leetcode

### c++

#### 依赖

gcc,g++,gdb,vscode

#### 编译/Compile

- VScode

  clone代码仓库

  ```shell
  git clone git@github.com:yiGmMk/leetcode.git
  ```

  在Vscode打开代码目录，选择文件->打开文件夹,选中代码目录

- vscode 安装依赖插件

  c++建议安装 twxs.cmake + austin.code-gnu-global+visualstudioexptteam.vscodeintellicode

- 点击对应的cpp源文件F5开始调试

### golang

#### lint

- 使用github action + golangci-lint
- lint配置文件: .golangci.yml

### py

#### issue

##### 2022-10-10,在contos和ubuntu18.04上默认安装的python3.6.9调试python代码断点不能触发

1. vscode版本
  Version: 1.72.0 (user setup)
  Commit: 64bbfbf67ada9953918d72e1df2f4d8e537d340e
  Date: 2022-10-04T23:20:39.912Z
  Electron: 19.0.17
  Chromium: 102.0.5005.167
  Node.js: 16.14.2
  V8: 10.2.154.15-electron.0
  OS: Windows_NT x64 10.0.19044
  Sandboxed: No

2. 升级python版本到3.9.12就好了,这里使用conda安装方便切换
   - 清华镜像 <https://mirrors.tuna.tsinghua.edu.cn/help/anaconda/>
   - miniconda <https://docs.conda.io/en/latest/miniconda.html>
   - linux下载的都是脚本通过脚本安装

#### fish 部分v2的版本支持conda

1. fish version(contos v2.3.1)
2. 报错: Variables may not be used as commands. In fish, please define a function or use 'eval $CONDA_EXE'.
3. 解决: 将v2的fish卸载,重新安装v3版本,再初始化conda的fish配置

   ```sh
   # 卸载(contos)
   yum erase fish

   # 安装
   cd /etc/yum.repos.d/
   wget https://download.opensuse.org/repositories/shells:fish:release:3/CentOS_7/shells:fish:release:3.repo
   yum install fish
   
   # 配置初始化
   conda init fish

   # 进入fish
   fish

   # 尝试使用
   conda -h

   # ubuntu 安装fish v3
   sudo apt-add-repository ppa:fish-shell/release-3
   sudo apt update
   sudo apt install fish
   ```

4. 参考 <https://github.com/conda/conda/issues/11079>
5. fish配置,输入fish_config,会弹出web页面,可视化配置
  ![fish_config](image/fish_config.png)
