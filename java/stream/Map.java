import java.util.stream.*;

/**
 * Map
 */
public class Map {
    // int stream
    public static Stream<Integer> getIntStream() {
        return Stream.of(1, 2, 3);
    }

    public static void main(String[] args) {
        getIntStream().map(n -> n * n).forEach(System.out::println);

        Stream.of("  Apple ", " pear ", " ORANGE", " BaNaNa ")
                .map(String::trim) // 去空格
                .map(String::toLowerCase) // 变小写
                .forEach(System.out::println); // 打印
    }
}