package base;

import java.util.Scanner;

public class IO {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.println("输入任意字符");
        String str = scanner.nextLine();
        System.out.println("输入任意数字");
        double val = scanner.nextDouble();
        System.out.println("字符:" + str + " 数值:" + val);
        scanner.close();
    }
}