
public class Solution {
    public static void main(String[] args) {
        AuthenticationManager m = new AuthenticationManager(5);
        m.renew("aaa", 1);
        m.generate("aaa", 2);

        System.out.println("6.cnt=" + m.countUnexpiredTokens(6));
        m.generate("bbb", 7);
        m.renew("aaa", 8);
        m.renew("bbb", 10);

        System.out.println("15.cnt=" + m.countUnexpiredTokens(15));
    }
}
