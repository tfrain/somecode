package structure

type LFUCache struct {
	// key 到 val 的映射，我们后文称为 KV 表
	keyToVal map[int]int
	// key 到 freq 的映射，我们后文称为 KF 表
	keyToFreq map[int]int
	// freq 到 key 列表的映射，我们后文称为 FK 表
	freqToKeys map[int]*linkedHashSet
	// 记录最小的频次
	minFreq int
	// 记录 LFU 缓存的最大容量
	cap int
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		keyToVal:   make(map[int]int),
		keyToFreq:  make(map[int]int),
		freqToKeys: make(map[int]*linkedHashSet),
		cap:        capacity,
		minFreq:    0,
	}
}

func (l *LFUCache) Get(key int) int {
	if _, ok := l.keyToVal[key]; !ok {
		return -1
	}
	l.increaseFreq(key)
	return l.keyToVal[key]
}

func (l *LFUCache) Put(key, value int) {
	// 维持健壮性?
	if l.cap <= 0 {
		return
	}
	/* 若 key 已存在，修改对应的 val 即可 */
	if _, ok := l.keyToVal[key]; ok {
		l.keyToVal[key] = value
		// key 对应的 freq 加一
		l.increaseFreq(key)
		return
	}
	/* key 不存在，需要插入 */
	/* 容量已满的话需要淘汰一个 freq 最小的 key */
	if l.cap <= len(l.keyToVal) {
		l.removeMinFreqKey()
	}
	/* 插入 key 和 val，对应的 freq 为 1 */
	// 插入 KV 表
	l.keyToVal[key] = value
	// 插入 KF 表
	l.keyToFreq[key] = 1
	// 插入 FK 表
	l.freqToKeys[1].add(key)
	// 插入新 key 后最小的 freq 肯定是 1
	l.minFreq = 1
}

func (l *LFUCache) increaseFreq(key int) {
	freq := l.keyToFreq[key]
	/* 更新 KF 表 */
	l.keyToFreq[key] = freq + 1
	/* 更新 FK 表 */
	// 将 key 从 freq 对应的列表中删除
	l.freqToKeys[freq].remove(key)
	// 将 key 加入 freq + 1 对应的列表中
	if l.freqToKeys[freq+1] == nil {
		l.freqToKeys[freq+1] = newLinkedHashSet()
	}
	l.freqToKeys[freq+1].add(key)
	// 如果 freq 对应的列表空了，移除这个 freq
	if l.freqToKeys[freq].size() == 0 {
		delete(l.freqToKeys, freq)
		// 如果这个 freq 恰好是 minFreq，更新 minFreq
		if freq == l.minFreq {
			l.minFreq++
		}
	}
}

func (l *LFUCache) removeMinFreqKey() {
	// freq 最小的 key 列表
	keyList := l.freqToKeys[l.minFreq]
	// 其中最先被插入的那个 key 就是该被淘汰的 key
	deletedKey := keyList.iterator().next()
	/* 更新 FK 表 */
	keyList.remove(deletedKey)
	if keyList.size() == 0 {
		delete(l.freqToKeys, l.minFreq)
		// 问：这里需要更新 minFreq 的值吗？
	}
	/* 更新 KV 表 */
	delete(l.keyToVal, deletedKey)
	/* 更新 KF 表 */
	delete(l.keyToFreq, deletedKey)
}

// 封装一个双向链表
type linkedHashSet struct {
	m    map[int]*node
	head *node
	tail *node
}

func newLinkedHashSet() *linkedHashSet {
	head, tail := new(node), new(node)
	head.next = tail
	tail.prev = head
	return &linkedHashSet{
		m:    make(map[int]*node),
		head: head,
		tail: tail,
	}
}

func (this *linkedHashSet) size() int {
	return len(this.m)
}

func (this *linkedHashSet) add(key int) {
	if _, ok := this.m[key]; ok {
		return
	}
	n := &node{key: key}
	last := this.tail.prev
	last.next = n
	n.prev = last
	n.next = this.tail
	this.tail.prev = n
	this.m[key] = n
}

func (this *linkedHashSet) remove(key int) {
	if n, ok := this.m[key]; ok {
		n.prev.next = n.next
		n.next.prev = n.prev
		delete(this.m, key)
	}
}

func (this *linkedHashSet) iterator() *keyIterator {
	return &keyIterator{this.head.next}
}

type node struct {
	key  int
	prev *node
	next *node
}

type keyIterator struct {
	n *node
}

func (this *keyIterator) hasNext() bool {
	return this.n.next != nil
}

func (this *keyIterator) next() int {
	this.n = this.n.next
	return this.n.key
}
