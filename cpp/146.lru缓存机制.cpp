/*
 * @lc app=leetcode.cn id=146 lang=cpp
 *
 * [146] LRU缓存机制
 */

// @lc code=start
// STL 已经为我们提供了 list 了，可以直接用起来。因为 key 是 int，使用 unordered_map 可能会稍快，相比 map 不太明显。
// hash 的 key 就是正常的 key，value 保存 list 中节点的地址。为什么要保存地址呢？因为方便将 list 中的节点进行移动。
// 你可以自己实现 list，这样的好处是你不用释放旧节点，就可以把旧节点直接移动到链表头，但是 c++ 的 std::list 似乎不支持将某个节点摘下来移动到头部，因此只能先释放掉，再插入新的头节点。虽然速度上比移动的办法要慢，但是我们实现 LRU 的思路和宗旨是不变的。

class LRUCache {
public:
    LRUCache(int capacity) : _cap(capacity) {}
    
    // O(1)
    // hash 查找，如果找到了，就把 list 中的节点接下来移到头部
    int get(int key) {
        auto it = _m.find(key);
        if (it == _m.end()) return -1;
        int val = it->second->second;
        _list.erase(it->second);
        _list.push_front(make_pair(key, val));
        _m[key] = _list.begin();
        return it->second->second;
    }
    
    // O(1)
    // 先查找旧 key 是否存在，如果存在，将节点移动到首部。
    // 如果不存在，插入新节点。
    // 如果容量超限，删除最脏的节点。
    void put(int key, int value) {
        auto it = _m.find(key);
        if (it != _m.end()) {
            _list.erase(it->second);
        }
        
        _list.push_front(make_pair(key, value));
        _m[key] = _list.begin();
        
        if (_list.size() > _cap) {
            int key = _list.back().first;
            _m.erase(key);
            _list.pop_back();
        }
    }
private:
    unordered_map<int, list<pair<int,int>>::iterator> _m;
    // 新节点或刚访问的节点插入表头，因为表头指针可以通过 begin 很方便的获取到。
    list<pair<int,int>> _list;
    int _cap;
};


struct DLinkedNode {
    int key, value;
    DLinkedNode* prev;
    DLinkedNode* next;
    DLinkedNode(): key(0), value(0), prev(nullptr), next(nullptr) {}
    DLinkedNode(int _key, int _value): key(_key), value(_value), prev(nullptr), next(nullptr) {}
};

class LRUCacheLinkNode {
private:
    unordered_map<int, DLinkedNode*> cache;
    DLinkedNode* head;
    DLinkedNode* tail;
    int size;
    int capacity;

public:
    LRUCache(int _capacity): capacity(_capacity), size(0) {
        // 使用伪头部和伪尾部节点
        head = new DLinkedNode();
        tail = new DLinkedNode();
        head->next = tail;
        tail->prev = head;
    }
    
    int get(int key) {
        if (!cache.count(key)) {
            return -1;
        }
        // 如果 key 存在，先通过哈希表定位，再移到头部
        DLinkedNode* node = cache[key];
        moveToHead(node);
        return node->value;
    }
    
    void put(int key, int value) {
        if (!cache.count(key)) {
            // 如果 key 不存在，创建一个新的节点
            DLinkedNode* node = new DLinkedNode(key, value);
            // 添加进哈希表
            cache[key] = node;
            // 添加至双向链表的头部
            addToHead(node);
            ++size;
            if (size > capacity) {
                // 如果超出容量，删除双向链表的尾部节点
                DLinkedNode* removed = removeTail();
                // 删除哈希表中对应的项
                cache.erase(removed->key);
                // 防止内存泄漏
                delete removed;
                --size;
            }
        }
        else {
            // 如果 key 存在，先通过哈希表定位，再修改 value，并移到头部
            DLinkedNode* node = cache[key];
            node->value = value;
            moveToHead(node);
        }
    }

    void addToHead(DLinkedNode* node) {
        node->prev = head;
        node->next = head->next;
        head->next->prev = node;
        head->next = node;
    }
    
    void removeNode(DLinkedNode* node) {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }

    void moveToHead(DLinkedNode* node) {
        removeNode(node);
        addToHead(node);
    }

    DLinkedNode* removeTail() {
        DLinkedNode* node = tail->prev;
        removeNode(node);
        return node;
    }
};
/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
// @lc code=end

