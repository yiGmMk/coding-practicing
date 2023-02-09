# IoC/Inversion of Control(控制反转)

IoC又称为依赖注入（DI：Dependency Injection）,它解决了
一个最主要的问题：将组件的创建+配置与组件的使用相分离，并且，
由IoC容器负责管理组件的生命周期

- [IoC/Inversion of Control(控制反转)](#iocinversion-of-control控制反转)
  - [IOC容器](#ioc容器)
  - [注解配置](#注解配置)
    - [Component](#component)
      - [生命周期](#生命周期)
    - [Autowired](#autowired)
    - [Configuration](#configuration)
    - [ComponentScan](#componentscan)
    - [Bean](#bean)
  - [Resource使用](#resource使用)
  - [条件注解](#条件注解)

## IOC容器

容器是一种为某种特定组件的运行提供必要支持的一个软件环境。例如，
Tomcat就是一个Servlet容器，它可以为Servlet的运行提供运行环境

Spring的核心就是提供了一个IoC容器，它可以管理所有轻量级的JavaBean组件，
提供的底层服务包括组件的生命周期管理、配置和组装服务、AOP支持

## 注解配置

```java
@Component
public class UserService {
    @Autowired
    MailService mailService;

    public UserService(@Autowired MailService mailService) {
        this.mailService = mailService;
    }
}
```

### Component

@Component注解相当于定义了一个Bean.

#### 生命周期

把一个Bean标记为@Component后，默认会自动为我们创建一个单例（Singleton），
即容器初始化时创建Bean，容器关闭前销毁Bean。在容器运行期间，我们调用getBean(Class)
获取到的Bean总是同一个实例。

### Autowired

@Autowired就相当于把指定类型的Bean注入到指定的字段中.

一般把@Autowired写在字段上，通常**使用package权限的字段**，便于测试

### Configuration

@Configuration表示是一个配置类

### ComponentScan

@ComponentScan告诉容器自动搜索当前类所在的包及子包,把所有标注为
@Component的Bean自动创建出来，并根据@Autowired进行装配

### Bean

[Bean](./Bean.md)

## Resource使用

使用Spring容器时，我们也可以把“文件”注入进来，方便程序读取,[Resource](./Resource.md)

## 条件注解

- @Profile
- @Conditional
- Springboot提供的@ConditionalOnProperty
