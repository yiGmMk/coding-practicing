package base;

public class Type {
    public static void main(String[] args) {
        Integer a = Integer.valueOf(111111111);
        Integer b = Integer.valueOf(111111111);
        System.out.println("a==b  :" + (a == b));
        System.out.println("equals:" + (a.equals(b)));
    }
}
