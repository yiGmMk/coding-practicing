import java.time.DayOfWeek;
import java.time.LocalDate;
import java.util.function.Supplier;
import java.util.stream.IntStream;
import java.util.stream.Stream;

/*
 * Stream.filter()是Stream的另一个常用转换方法。
 * 所谓filter()操作，就是对一个Stream的所有元素一一进行测试，
 * 不满足条件的就被“滤掉”了，剩下的满足条件的元素就构成了一个新的Stream
*/
public class Filter {
    public static void main(String[] args) {
        System.out.println("filter number");
        IntStream.of(1, 2, 3, 4, 5).filter(n -> n >= 3)
                .forEach(System.out::println);

        System.out.println("filter date");
        Stream.generate(new LocalDateSupplier())
                .limit(31)
                .filter(ldt -> ldt.getDayOfWeek() == DayOfWeek.SATURDAY ||
                        ldt.getDayOfWeek() == DayOfWeek.SUNDAY)
                .forEach(System.out::println);
    }
}

class LocalDateSupplier implements Supplier<LocalDate> {
    LocalDate start = LocalDate.of(2023, 1, 1);
    int n = -1;

    public LocalDate get() {
        n++;
        return start.plusDays(n);
    }
}