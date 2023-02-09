# Resource

注入Resource最常用的方式是通过classpath，即类似classpath:/logo.txt
表示在classpath中搜索logo.txt文件

使用Maven的标准目录结构，所有资源文件放入src/main/resources即可

```java
@Component
public class AppService {
    @Value("classpath:/logo.txt")
    private Resource resource;

    private String logo;

    @PostConstruct
    public void init() throws IOException {
        try (var reader = new BufferedReader(
                new InputStreamReader(resource.getInputStream(), StandardCharsets.UTF_8))) {
            this.logo = reader.lines().collect(Collectors.joining("\n"));
        }
    }
}
```
