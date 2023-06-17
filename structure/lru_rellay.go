package structure

import (
	"container/list"
	"sync"
)

type item struct {
	key   string
	value interface{}
}

type LRU struct {
	capacity  int
	itemsList *list.List
	itemsMap  map[string]*list.Element
	lock      sync.Mutex
}

func New(capacity int) *LRU {
	return &LRU{
		capacity:  capacity,
		itemsList: list.New(),
		itemsMap:  make(map[string]*list.Element),
	}
}

func (l *LRU) Put(key string, cacheItem interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if node, ok := l.itemsMap[key]; ok {
		l.itemsList.MoveToFront(node)
		return
	}
	if l.capacity == len(l.itemsMap) {
		delete(l.itemsMap, l.itemsList.Back().Value.(item).key)
		l.itemsList.Remove(l.itemsList.Back())
	}
	node := l.itemsList.PushFront(item{key: key, value: cacheItem})
	l.itemsMap[key] = node
}

func (l *LRU) Get(key string) interface{} {
	if node, ok := l.itemsMap[key]; ok {
		l.itemsList.MoveToFront(node)
		return node.Value.(item).value
	}
	value := getValueFromDB(key)
	l.Put(key, value)

	return value
}

func (l *LRU) Size() int {
	return l.itemsList.Len()
}

func (l *LRU) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.itemsList = list.New()
	l.itemsMap = make(map[string]*list.Element)
}

func getValueFromDB(key string) interface{} {
	value := key
	return value
}
