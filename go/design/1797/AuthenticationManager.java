import java.util.Comparator;
import java.util.HashMap;
import java.util.PriorityQueue;
import java.util.stream.Stream;

public class AuthenticationManager {
    int t;
    PriorityQueue<Token> q;
    HashMap<String, Token> m;

    public AuthenticationManager(int timeToLive) {
        this.t = timeToLive;
        this.q = new PriorityQueue<>();
        this.m = new HashMap<String, Token>();
    }

    public void remove(int currentTime) {
        if (currentTime < t) {
            return;
        }
        currentTime = currentTime - t;
        while (q.size() > 0) {
            Token t = q.peek();
            if (t.t <= currentTime) {
                q.remove(t);
                m.remove(t.token);
            } else {
                break;
            }
        }
    }

    public void generate(String tokenId, int currentTime) {
        remove(currentTime);

        Token t = new Token(tokenId, currentTime);
        q.add(t);
        m.put(tokenId, t);
    }

    public void renew(String tokenId, int currentTime) {
        remove(currentTime);
        if (!m.containsKey(tokenId)) {
            return;
        }
        Token t = m.get(tokenId);
        q.remove(t);

        t.t = currentTime;
        m.replace(tokenId, t);
        q.offer(t);
    }

    public int countUnexpiredTokens(int currentTime) {

        remove(currentTime);

        return q.size();
    }

    public static void main(String[] args) {
        // 要使用PriorityQueue,要么Token实现Comparable<Token>接口(compareTo方法)
        // 要是实现一个TokenComparator(override compare方法)
        PriorityQueue<Token> q = new PriorityQueue<>(); // new PriorityQueue<>(new TokenComparator());
        q.offer(new Token("a", 0));
        q.offer(new Token("a", 10));
        q.offer(new Token("a", 3));
        q.offer(new Token("c", 2));
        q.offer(new Token("a", 2));
        q.offer(new Token("b", 2));

        // poll 会移除元素,第二次打印没有数据了
        Stream.generate(q::poll).limit(q.size()).forEach(System.out::println);
        System.out.println("-----2----");
        Stream.generate(q::poll).limit(q.size()).forEach(System.out::println);
    }
}

class TokenComparator implements Comparator<Token> {
    // Compares its two arguments for order. Returns a negative integer, zero, or a
    // positive integer as the first argument is less than, equal to, or greater
    // than the second.
    @Override
    public int compare(Token o1, Token o2) {
        if (o1.t < o2.t) {
            return -1;
        } else if (o1.t == o2.t) {
            return o1.token.compareTo(o2.token);
        }
        return 1;
    }
}

class Token implements Comparator<Token>, Comparable<Token> {
    public String token;
    public int t;

    public Token(String s, int t) {
        this.token = s;
        this.t = t;
    }

    public String toString() {
        return "token:" + token + "," + t;
    }

    @Override
    public int compareTo(Token o2) {
        if (t < o2.t) {
            return -1;
        } else if (t == o2.t) {
            return token.compareTo(o2.token);
        }
        return 1;
    }

    @Override
    public int compare(Token o1, Token o2) {
        if (o1.t < o2.t) {
            return -1;
        } else if (o1.t == o2.t) {
            return o1.token.compareTo(o2.token);
        }
        return 1;
    }
}
