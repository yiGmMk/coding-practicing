# design pattern

设计模式是软件设计中常见问题的典型解决方案

![all](./view.PNG)

## 创建型模式/creational-patterns

### 工厂模式

工厂模式包含简单工厂、工厂方法、抽象工厂、DI容器

> 在 GoF 的《设计模式》一书中，它将简单工厂模式看作是工厂方法模式的一种特例，所以工厂模式只被分成了工厂方法和抽象工厂两类。

1. DI/IOC(依赖注入/控制翻转)

> 准备一个盒子(容器),事先将项目中可能用到的类扔进去,在项目中直接从容器中拿,也就是避免了直接在项目中到处new,造成大量耦合。取而代之的是在项目类里面增设 setDi()和getDi()方法,通过Di同一管理类

- DI(dependency injection)
- 使用generate实现,go工具包wire(Google 开源的一个依赖注入工具,它是一个代码生成器,并不是一个框架)  <https://github.com/google/wire>
- 使用反射实现的,<https://github.com/uber-go/dig>

### 工厂方法/Factory Method

工厂方法模式在父类中提供一个创建对象的方法,允许子类决定实例化对象的类型

### 抽象工厂/Abstract Factory

能**创建一系列相关的对象， 而无需指定其具体类**。
在许多设计工作的初期都会使用工厂方法模式 （较为简单， 而且可以更方便地通过子类进行定制）， 随后演化为使用抽象工厂模式、 原型模式或生成器模式 （更灵活但更加复杂）。

### 生成器/Build

### 原型/Prototype

### 单例/Singleton

1. 饿汉式
   不管用不用,先创建
2. 懒汉式
   使用时才创建

## 结构型模式

## 行为模式

## references

- <https://refactoring.guru/design-patterns/what-is-pattern>
- code [git@github.com:RefactoringGuru/design-patterns-go.git]<https://github.com/RefactoringGuru/design-patterns-go>
- [git@github.com:mohuishou/go-design-pattern.git](https://github.com/mohuishou/go-design-pattern)
- di容器 <https://www.cnblogs.com/feichengwulai/articles/6266555.html>
- 极客时间 <https://time.geekbang.org/column/article/197254>
- [设计模式之美]<https://weread.qq.com/web/bookDetail/1fb322f0811e7c245g013d1c>
- <https://toutiao.io/k/bjjvhwf>
