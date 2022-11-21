# coredump

core dump文件是在程序异常终止或崩溃时操作系统对程序当时的内存状态进行记录并保存而生成的一个数据文件，该文件以core命名，也被称为核心转储文件

go程序崩溃默认不会产生升core文件,需要进行设置.

## 设置

不限制core文件的大小

```bash
ulimit -c unlimited
```

启动参数

```bash
GOTRACEBACK=crash ./xxx
```

## 使用delve调试core文件

```bash
#Usage:
#dlv core <executable> <core> [flags]
dlv core ./main ./core.xxx 
```
