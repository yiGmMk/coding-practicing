# java

- [java](#java)
  - [类型](#类型)
    - [包装类型](#包装类型)
      - [包装类对比](#包装类对比)
  - [枚举](#枚举)
  - [JavaBean](#javabean)
    - [枚举JavaBean属性](#枚举javabean属性)
  - [记录类](#记录类)
  - [日志](#日志)
  - [包管理](#包管理)
  - [参考](#参考)

## 类型

Java的数据类型分两种：

基本类型：byte，short，int，long，boolean，float，double，char

引用类型：所有class和interface类型

引用类型可以赋值为null，表示空，但基本类型不能赋值为null

### 包装类型

将基本类型类型用类封装

Java核心库为每种基本类型都提供了对应的包装类型,如java.lang.Boolean,java.lang.Double

核心库提供的包装类**都是不变类(final)**,创建后是不变的

```java
@jdk.internal.ValueBased
public final class Integer extends Number
        implements Comparable<Integer>, Constable, ConstantDesc {
    private final int value;        
}
```

#### 包装类对比

核心库提供的包装类实例必须使用equals对比不能使用==

```java
 Integer a = Integer.valueOf(111111111);
 Integer b = Integer.valueOf(111111111);
 System.out.println("a==b  :" + (a == b));
 System.out.println("equals:" + (a.equals(b)));
```

## 枚举

enum定义的类型是class,有以下几个特点：

- 定义的enum类型总是继承自java.lang.Enum，且无法被继承；
- 只能定义出enum的实例，而无法通过new操作符创建enum的实例；
- 定义的每个实例都是引用类型的唯一实例；
- 可以将enum类型用于switch语句。

```java
enum Weekday {
    MON(1, "星期一"), TUE(2, "星期二"), WED(3, "星期三"), THU(4, "星期四"), FRI(5, "星期五"), SAT(6, "星期六"), SUN(0, "星期日");

    public final int dayValue;
    private final String chinese;

    private Weekday(int dayValue, String chinese) {
        this.dayValue = dayValue;
        this.chinese = chinese;
    }

    // 默认情况下，对枚举常量调用toString()会返回和name()一样的字符串
    // toString()可以被覆写，而name()则不行
    @Override
    public String toString() {
        return this.chinese;
    }
}
```

## JavaBean

读写方法符合以下这种命名规范的class被称为JavaBean

```java
// 读方法:
public Type getXyz()
// 写方法:
public void setXyz(Type value)
```

JavaBean主要用来传递数据，即把一组数据组合成一个JavaBean便于传输。
此外，JavaBean可以方便地被IDE工具分析，生成读写属性的代码，主要用在图形界面的可视化设计中。

### 枚举JavaBean属性

可以使用Java核心库提供的Introspector枚举JavaBean的所有属性

```java
BeanInfo info = Introspector.getBeanInfo(Person.class);
for (PropertyDescriptor pd : info.getPropertyDescriptors()) {
    System.out.println(pd.getName());
    System.out.println("  " + pd.getReadMethod());
    System.out.println("  " + pd.getWriteMethod());
}
// 输出
// age
//   public int base.Person.getAge()
//   public void base.Person.setAge(int)
// class
//   public final native java.lang.Class java.lang.Object.getClass()
//   null
// name
//   public java.lang.String base.Person.getName()
//   public void base.Person.setName(java.lang.String)
```

其中"class"是从Object继承的getClass()方法带来的

## 记录类

从Java 14开始，引入了新的Record类

```java
public class Main {
    public static void main(String[] args) {
        Point p = new Point(123, 456);
        System.out.println(p.x());
        System.out.println(p.y());
        System.out.println(p);
    }
}

public record Point(int x, int y) {}
```

除了用final修饰class以及每个字段外，编译器还自动为我们创建了构造方法，和字段名同名的方法，以及覆写toString()、equals()和hashCode()方法。

换句话说，使用record关键字，可以一行写出一个不变类。

## 日志

Commons Logging和Log4j: 一个负责充当日志API，一个负责实现日志底层，搭配使用非常方便

## 包管理

## 参考

- <https://www.liaoxuefeng.com/wiki/1252599548343744/1331429187256353>
