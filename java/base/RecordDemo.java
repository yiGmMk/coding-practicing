package base;

public class RecordDemo {
    public static void main(String[] args) {
        Point p = new Point(123, 456);
        System.out.println(p.x());
        System.out.println(p.y());
        System.out.println(p);

    }
}

record Point(int x, int y) {
}