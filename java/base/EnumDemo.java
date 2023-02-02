package base;

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

public class EnumDemo {
    public static void main(String[] args) {
        Weekday monday = Weekday.MON;
        System.out.println("Today is " + monday + ". Work at office!");
    }
}
