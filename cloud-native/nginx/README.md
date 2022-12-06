# nignx

## root,location,alias

### root与alias

假如服务器路径为：/home/imooc/files/img/face.png

- root路径完全匹配访问

 配置的时候为:

 locaion /imooc{
     root /home
 }
 用户访问的时候请求为：url:port/imooc/files/img/face.png

- alias可以为你的路径做一个别名，对用户透明配置的时候为：

 location /hello{
     alias /home/imooc
 }
 用户访问的时候请求为：url:port/hello/files/img/face.png,如此相当于目录imooc做一个自定义的别名。

### location的匹配规则

#### 空格：默认匹配，普通匹配

localtion /{
    root /home;
}

#### =:精确匹配

location = /imooc/img/face1.png {
    root /home;
}

#### ~*:匹配正则表达式，不区分大小写

```shell
# 符合图片的显示
location ~* .(GIT|jpg|png|jpeg) {
    root /home;
}
```

#### ~:匹配正则表达式，区分大小写

```sh
# GIF必须大写才能匹配到
location ~ .(GIF|jpg|png|jpeg) {
    root /home;
}
```

#### ^~:以某个字符路径开头

```sh
location ^~ /imooc/img {
    root /home;
}
```

### 参考

- <http://t.csdn.cn/J3npY>
- book: NGINX Cookbook(2022 O'Reilly Media)

## Authentication

支持basic authentication,authentication subrequests(子请求认证)
NGINX PLUS还额外支持JWT(JSON Web Tokens)

### basic authentication

basic auth需要以以下格式存储认证信息,以冒号分隔开用户名,密码,备注等,其中密码需要加密或使用哈希值.

```conf
# comment
name1:password1
name2:password2:comment
name3:password3
```

nginx配置

```conf
location / {
    auth_basic          "Private site";
    auth_basic_user_file conf.d/passwd; # 认证信息文件
}
```

在http, server, or location里都支持basic authentication

#### 加密

nginx仅支持部分加密算法,如openssl的passwd命令的默认加密方法

```sh
# 默认为 -crypt ,等同于 openssl passwd -crypt abcde
openssl passwd abcde
# Ek3zv167GN/Ow

#以表格形式输出
openssl passwd -table -crypt abcde
#abcde   BbdA4jamildcM
```

### authentication subrequests
