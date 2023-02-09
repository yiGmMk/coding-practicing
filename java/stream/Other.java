import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

/**
 * Other
 */
public class Other {
    public static void main(String[] args) {
        // 排序
        System.out.println("sort1");
        Stream.of("B", "a", "D", "c").sorted()
                .collect(Collectors.toList()).forEach(System.out::println);

        System.out.println("sort2.忽略大小写");
        Stream.of("B", "a", "D", "c")
                .sorted(String::compareToIgnoreCase)
                .collect(Collectors.toList()).forEach(System.out::println);

        // 截取
        System.out.println("sort2.忽略大小写+skip+limit");
        Stream.of("B", "a", "D", "c")
                .sorted(String::compareToIgnoreCase)
                .limit(2)
                .skip(1)
                .collect(Collectors.toList()).forEach(System.out::println);

        // 合并
        System.out.println("sort2.忽略大小写+concat");
        Stream.concat(
                Stream.of("B", "a", "D", "c").sorted(String::compareToIgnoreCase),
                Stream.of("B", "a", "D", "c").sorted(String::compareToIgnoreCase))
                .collect(Collectors.toList())
                .forEach(System.out::println);

        // flatMap
        Stream<List<Integer>> s = Stream.of(
                Arrays.asList(1, 2, 3, 10),
                Arrays.asList(4, 5, 6),
                Arrays.asList(7, 8, 9));
        Stream<Integer> i = s.flatMap(list -> list.stream());
        System.out.println("flatMap");
        i.forEach(System.out::println);

        // 并行
    }
}