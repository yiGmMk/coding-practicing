package base;

import java.util.Arrays;
import java.util.Comparator;

/**
 * FunctionalProgrammingDemo
 * 函数式编程
 */
public class FunctionalProgrammingDemo {
    static int cmp(String s1, String s2) {
        return s1.compareTo(s2);
    }

    // 从Java 8开始，我们可以用Lambda表达式替换单方法接口
    void SortStrings(String[] strs) {
        Arrays.sort(strs, new Comparator<String>() {
            public int compare(String s1, String s2) {
                return s1.compareTo(s2);
            }
        });
    }

    /*
     * Lambda
     * 参数是(s1, s2)，参数类型可以省略，因为编译器可以自动推断出String类型。-> { ...}
     * 表示方法体，所有代码写在内部即可。Lambda表达式没有class定义，因此写法非常简洁。
     */
    void LambdaSortStrings(String[] strs) {
        Arrays.sort(strs, (s1, s2) -> {
            return s1.compareTo(s2);
        });
        // 等同于 Arrays.sort(array, (s1, s2) -> s1.compareTo(s2));
    }

    // 还可以引入签名一致的方法
    // 这里,方法签名只看参数类型和返回类型，不看方法名称，也不看类的继承关系
    public static void main(String[] args) {
        String[] array = new String[] { "Apple", "Orange", "Banana", "Lemon" };
        Arrays.sort(array, FunctionalProgrammingDemo::cmp);
        System.out.println(String.join(", ", array));

        String[] array1 = new String[] { "Apple", "Orange", "Banana", "Lemon" };
        Arrays.sort(array1, String::compareTo);
        System.out.println(String.join(", ", array));
    }
}