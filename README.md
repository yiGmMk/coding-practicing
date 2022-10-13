# coding & practicing

## design patterns

![design patterns](./design-pattern/view.PNG)

### references

- [refactoringguru.cn](https://refactoringguru.cn/design-patterns/abstract-factory)
- [博文,设计模式](https://lailin.xyz/post/factory.html)
- [博文,刘丹冰aceld,设计模式](https://www.yuque.com/aceld/lfhu8y/pebesh)

## algorithm & datastructure

### algorithm

#### sort

#### search

#### bitmap

### datastructure

#### tree

#### list

#### array

## awesome xxx & mark

### collection

1. [awesome-go-cn,Go框架、库和软件的中文收录大全](https://github.com/yinggaozhen/awesome-go-cn)
2. [开发者头条,博客/公众号各种站点资讯,提供订阅服务](https://toutiao.io/posts/hot/30)
3. [架构师](https://github.com/xingshaocheng/architect-awesome)
4. [github 资讯](https://github.com/GitHubDaily/GitHubDaily)

### database

1. [SQL 审核查询平台](https://github.com/hhyo/Archery)
2. [极客时间 丁奇 Mysql45讲](https://time.geekbang.org/column/intro/100020801)
3. [mysql 日期相关函数](https://blog.csdn.net/hu1010037197/article/details/115391335)
4. [博文:LeetCode题解 mysql 查询策略,窗口函数,排序](https://leetcode.cn/problems/nth-highest-salary/solution/mysql-zi-ding-yi-bian-liang-by-luanz/)
5. [知乎专栏,mysql 优化](https://www.zhihu.com/column/c_1439702957097447424)
6. [MySQL TIME_WAIT连接,2MSL](https://blog.csdn.net/fkew2009/article/details/86714699)

### blog

1. [小林coding](https://xiaolincoding.com/)

### notes

- [笔记,汇总 programnotes.cn](https://programnotes.cn/a-set-of-notes-blogs/index.html)

### linux

1. [shell配置,bashrc,profile](https://zhuanlan.zhihu.com/p/405174594)

### language

- [python](https://docs.python.org/zh-cn/3.10/tutorial/index.html)

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

### python

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
