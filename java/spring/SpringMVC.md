# Spring MVC

Java Web的基础：Servlet容器，以及标准的Servlet组件

## CORS

[CORS](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CORS)，全称Cross-Origin Resource Sharing，是HTML5规范定义的如何跨域访问资源。
默认情况下，浏览器按同源策略放行JavaScript调用API，即:

- 如果A站在域名a.com页面的JavaScript调用A站自己的API时，没有问题；
- 如果A站在域名a.com页面的JavaScript调用B站b.com的API时，将被浏览器拒绝访问，因为不满足同源策略。

如果A站的JavaScript访问B站API的时候，B站能够返回响应头Access-Control-Allow-Origin: <http://a.com，那么，浏览器就允许A站的JavaScript访问B站的API>。
注意到跨域访问能否成功，取决于B站是否愿意给A站返回一个正确的Access-Control-Allow-Origin响应头

### config

#### 使用@CrossOrigin注解

```java
@CrossOrigin(origins = "http://local.liaoxuefeng.com:8080")
@RestController
@RequestMapping("/api")
public class ApiController {
    ...
}
```

#### 在WebMvcConfigurer定义全局CORS配置

```java
@Bean
WebMvcConfigurer createWebMvcConfigurer() {
    return new WebMvcConfigurer() {
        @Override
        public void addCorsMappings(CorsRegistry registry) {
            registry.addMapping("/api/**")
                    .allowedOrigins("http://local.liaoxuefeng.com:8080")
                    .allowedMethods("GET", "POST")
                    .maxAge(3600);
            // 可以继续添加其他URL规则:
            // registry.addMapping("/rest/v2/**")...
        }
    };
}
```
