package structure

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

// https://github.com/charizer/lrucache/blob/master/lrucache.go
// EvictCallback is used to get a callback when a cache entry is evicted
type EvictCallback func(key string, value interface{})

type element struct {
	key    string
	value  interface{}
	expire *time.Time
}

func (e *element) IsExpired() bool {
	if e.expire == nil {
		return false
	}
	return time.Now().After(*e.expire)
}

type LRUExpire struct {
	capacity   int
	itemsMap   map[string]*list.Element
	itemList   *list.List
	onEvict    EvictCallback
	defaultTTL time.Duration
	mu         sync.Mutex
}

func NewLRUExpire(capacity int) (*LRUExpire, error) {
	if capacity < 0 {
		return nil, errors.New("capicity must be positive")
	}
	return &LRUExpire{
		capacity: capacity,
		itemsMap: make(map[string]*list.Element),
		itemList: list.New(),
	}, nil
}

func (l *LRUExpire) Put(key string, value interface{}, ttl time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()
	var expire time.Time
	if ttl > 0 {
		expire = time.Now().Add(ttl)
	} else if l.defaultTTL > 0 {
		expire = time.Now().Add(l.defaultTTL)
	}
	if v, ok := l.itemsMap[key]; ok {
		l.itemList.MoveToFront(v)
		// update time
		v.Value.(*element).value = value
		v.Value.(*element).expire = &expire
		return
	}
	if len(l.itemsMap) == l.capacity {
		delete(l.itemsMap, l.itemList.Back().Value.(*element).key)
		l.itemList.Remove(l.itemList.Back())
	}
	node := l.itemList.PushFront(&element{
		key:    key,
		value:  value,
		expire: &expire,
	})
	l.itemsMap[key] = node
}

func (l *LRUExpire) Get(key string) interface{} {
	if v, ok := l.itemsMap[key]; ok {
		if v.Value.(*element).IsExpired() {
			l.removeElement(v)
			return nil
		}
		l.itemList.MoveToFront(v)
		return v.Value.(*element).value
	}
	return nil
}

// removeElement is used to remove a given list element from the cache
func (l *LRUExpire) removeElement(e *list.Element) {
	l.itemList.Remove(e)
	node := e.Value.(*element)
	delete(l.itemsMap, node.key)
	if l.onEvict != nil {
		l.onEvict(node.key, node.value)
	}
}
