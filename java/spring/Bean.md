# Bean

## 定制Bean

```java
@Component
public class MailService {
    @Autowired(required = false)
    ZoneId zoneId = ZoneId.systemDefault();

    @PostConstruct
    public void init() {
        System.out.println("Init mail service with zoneId = " + this.zoneId);
    }

    @PreDestroy
    public void shutdown() {
        System.out.println("Shutdown mail service");
    }
}
```

### 初始化/清理

Bean的初始化和清理方法标记:@PostConstruct,@PreDestroy

### 别名

默认情况下，对一种类型的Bean，容器只创建一个实例,当需要对一种类型的Bean创建多个实例需要使用别名
否则Spring会报NoUniqueBeanDefinitionException异常，意思是出现了**重复的Bean定义**

可以用@Bean("name")指定别名，也可以用@Bean+@Qualifier("name")指定别名

```java
@Configuration
@ComponentScan
public class AppConfig {
    @Bean("z")
    ZoneId createZoneOfZ() {
        return ZoneId.of("Z");
    }

    @Bean
    @Qualifier("utc8")
    ZoneId createZoneOfUTC8() {
        return ZoneId.of("UTC+08:00");
    }
}


@Component
public class MailService {
 @Autowired(required = false)
 @Qualifier("z") // 指定注入名称为"z"的ZoneId
 ZoneId zoneId = ZoneId.systemDefault();
    ...
}
```

注入时也需要指定名称

### Prototype

Spring默认使用Singleton创建Bean，也可指定Scope为Prototype

每次调用getBean(Class)，容器都返回一个新的实例，这种Bean称为Prototype（原型），它的生命周期和Singleton不同

### 可选注入

用@Autowired(required=false)允许可选注入
