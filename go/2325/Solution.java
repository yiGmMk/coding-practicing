import java.util.HashMap;
import java.util.Map;

class Solution {
    public String decodeMessage(String key, String message) {
        char cur = 'a';
        Map<Character, Character> m = new HashMap<>();
        for (int i = 0; i < key.length(); i++) {
            if (key.charAt(i) == ' ') {
                continue;
            }
            if (!m.containsKey(key.charAt(i))) {
                m.put(key.charAt(i), cur);
                cur++;
            }
        }
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < message.length(); i++) {
            if (message.charAt(i) == ' ') {
                sb.append(' ');
            } else {
                sb.append(m.get(message.charAt(i)));
            }
        }
        return sb.toString();
    }
}
