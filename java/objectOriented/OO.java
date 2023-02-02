package objectOriented;

// 方法名相同，但各自的参数不同，称为方法重载（Overload）
// 在继承关系中，子类如果定义了一个与父类方法签名完全相同的方法，被称为覆写（Override）
// Override和Overload不同的是，如果方法签名不同，就是Overload，Overload方法是一个新方法；
// 如果方法签名相同，并且返回值也相同，就是Override
// override可以加上@Override标签
public class OO {
    // instanceof实际上判断一个变量所指向的实例是否是指定类型，或者这个类型的子类。
    // 如果一个引用变量为null，那么对任何instanceof的判断都为false
    public static void instanceUsage() {
        Object obj = "string object";
        // 从Java 14开始，判断instanceof后，可以直接转型为指定变量，避免再次强制转型
        if (obj instanceof String s) {
            System.out.println("obj:" + obj.toString() + " instanceof:" + s);
        }
    }

    public static void override() {
        System.out.println();

        Person s = new Student("王二", 100, 100);
        Person p = new Person("王二", 1);

        System.out.println("p_student:" + s.toString() + " p_person:" + p.toString());
        System.out.println("ps:" + s.getName());
        System.out.println("pp:" + p.getName());
    }

    public static void main(String[] args) {
        // instanceof使用
        System.out.println("instanceof usage");
        instanceUsage();

        System.out.println("override,overload usage");
        // 多态,重载,重写
        // Java的实例方法调用是基于运行时的实际类型的动态调用，而非变量的声明类型。
        // 这个非常重要的特性在面向对象编程中称之为多态 Polymorphic。
        override();
    }
}
