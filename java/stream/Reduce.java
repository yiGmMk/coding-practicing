
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Reduce {
    public static void main(String[] args) {
        // reduce()方法传入的对象是BinaryOperator接口，它定义了一个apply()方法，
        // 负责把上次累加的结果和本次的元素进行运算，并返回累加的结果
        int sum = Stream.of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
                .reduce(0, (acc, n) -> acc + n);
        // reduce()操作首先初始化结果为指定值（这里是0），紧接着，
        // reduce()对每个元素依次调用(acc, n) -> acc + n，其中，acc是上次计算的结果
        System.out.println("sum=" + sum);

        System.out.println("----map+reduce------------");
        Stream<String> stream = Stream.of("APPL:Apple", "MSFT:Microsoft");
        System.out.println(stream.collect(Collectors.toMap(
                // 把元素s映射为key:
                s -> s.substring(0, s.indexOf(':')),
                // 把元素s映射为value:
                s -> s.substring(s.indexOf(':') + 1))).toString());

        // Collectors.groupingBy()，它需要提供两个函数：一个是分组的key
        // 第二个是分组的value，这里直接使用Collectors.toList()，表示输出为List
        System.out.println(Stream.of("APPL:Apple", "MSFT:Microsoft", "APPL:IPAD", "APPLE:APONE")
                .collect(Collectors.groupingBy(
                        s -> s.substring(0, s.indexOf(':')),
                        // 把元素s映射为value:
                        Collectors.toList()))
                .toString());
    }
}
