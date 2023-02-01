class Solution:
    def decodeMessage(self, key: str, message: str) -> str:
        cur = 'a'
        m = dict()
        for c in key:
            if c == ' ':
                continue
            if c not in m:
                m[c] = cur
                cur = chr(ord(cur)+1)
        ans = "".join(m.get(c, " ") for c in message)
        return ans
