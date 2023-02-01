# object oriented programing 面向对象编程

- [object oriented programing 面向对象编程](#object-oriented-programing-面向对象编程)
  - [polymorphic 多态](#polymorphic-多态)
    - [override function from Object](#override-function-from-object)
    - [super](#super)
    - [final](#final)
  - [Abstract,抽象类](#abstract抽象类)
  - [Interface,接口](#interface接口)
    - [接口继承](#接口继承)
    - [default 方法](#default-方法)
  - [静态字段/静态方法](#静态字段静态方法)
  - [package 包](#package-包)
    - [包作用域](#包作用域)
  - [作用域](#作用域)
    - [private](#private)
    - [protected](#protected)
    - [package](#package)

## polymorphic 多态

多态是指，针对某个类型的方法调用，其真正执行的方法取决于运行时期实际类型的方法

```java
// 它传入的参数类型是Person，我们无法知道传入的参数实际类型究竟是Person，
// 还是Student，还是Person的其他子类，
// 因此，也无法确定调用的是不是Person类定义的run()方法
public void runTwice(Person p) {
    p.run();
    p.run();
}
```

多态的例子[计税](https://www.liaoxuefeng.com/wiki/1252599548343744/1260455778791232)
不同时收入计税方式不同,利用多态可以很方便地实现

### override function from Object

java中所有的类都是Object的子类,可以override Object的一些方法

- toString()：把instance输出为String；
- equals()：判断两个instance是否逻辑相等；
- hashCode()：计算一个instance的哈希值。

### super

在子类中可以通过super调用父类被override的方法

### final

如果一个父类不允许子类对它的某个方法进行覆写，可以把该方法标记为final。用final修饰的方法不能被Override

如果一个类不希望任何其他类继承自它，那么可以把这个类本身标记为final。用final修饰的类不能被继承

对于一个类的实例字段，同样可以用final修饰。用final修饰的字段在初始化后不能被修改

## Abstract,抽象类

```java
abstract class Person {
    public abstract void run();
}
// 把一个方法声明为abstract，表示它是一个抽象方法，
// 本身没有实现任何方法语句。因为这个抽象方法本身是无法执行的，
// 所以，Person类也无法被实例化。
// 编译器会告诉我们，无法编译Person类，因为它包含抽象方法。
```

无法实例化的抽象类有什么用？

因为抽象类本身被设计成只能用于被继承，因此，抽象类可以强迫子类实现其定义的抽象方法，否则编译会报错。因此，抽象方法实际上相当于定义了“规范”。

## Interface,接口

interface，就是比抽象类还要抽象的纯抽象接口，因为它连字段都不能有(除final类型的静态字段外)。
因为接口定义的所有方法默认都是public abstract的，所以这两个修饰符不需要写出来（写不写效果都一样）

```java
interface Person {
    void run();
    String getName();
}
```

在Java中，一个类只能继承自另一个类，不能从多个类继承。但是，一个类可以实现多个interface

```java
class Student implements Person, Hello { // 实现了两个interface
    ...
}
```

### 接口继承

一个interface可以继承自另一个interface。
interface继承自interface使用extends，它相当于扩展了接口的方法

一般来说，公共逻辑适合放在abstract class中，具体逻辑放到各个子类，而接口层次代表抽象程度。
可以参考Java的[集合类定义的一组接口、抽象类以及具体子类的继承关系](https://www.liaoxuefeng.com/wiki/1252599548343744/1260456790454816)

### default 方法

实现类可以不必override default方法。
default方法的目的是，当我们需要给接口新增一个方法时，会涉及到修改全部子类。
如果新增的是default方法，那么子类就不必全部修改，只需要在需要override的地方去override新增方法。

## 静态字段/静态方法

```java
class Person {
    public String name;
    public int age;
    // 定义静态字段number:
    public static int number;
}
```

推荐用类名来访问静态字段。可以把静态字段理解为描述class本身的字段（非实例字段）,所有实例共享静态字段

静态方法属于class而不属于实例，
因此，静态方法内部，无法访问this变量，也无法访问实例字段，它只能访问静态字段。

interface可以有 public static final类型的静态字段

```java
public interface Person {
    // 编译器会自动加上public statc final:
    int MALE = 1;
    int FEMALE = 2;
}
```

## package 包

在Java虚拟机执行的时候，JVM只看完整类名，因此，只要包名不同，类就不同

```java
package types; // 申明包名types
public class Arrays {
}

package util; // 申明包名util
public class Arrays {
}
```

### 包作用域

位于同一个包的类，可以访问包作用域的字段和方法。
不用public、protected、private修饰的字段和方法就是包作用域。

```java
package hello;

public class Person {
    // 包作用域:
    void hello() {
        System.out.println("Hello!");
    }
}
Main类也定义在hello包下面：

package hello;

public class Main {
    public static void main(String[] args) {
        Person p = new Person();
        p.hello(); // 可以调用，因为Main和Person在同一个包
    }
}
```

## 作用域

用来限定访问作用域的修饰符:public protected private

### private

定义为private的field、method无法被其他类访问,确切地说，private访问权限被限定在class的内部，而且与方法声明顺序无关

推荐把private方法放到后面，因为public方法定义了类对外提供的功能，阅读代码的时候，应该先关注public方法

### protected

protected作用于继承关系。定义为protected的字段和方法可以被子类访问，以及子类的子类

### package

包作用域是指一个类允许访问同一个package的

- 没有public、private修饰的class，
- 以及没有public、protected、private修饰的字段和方法
