package cachepkg

import (
	"sync"
	"time"
)

type SimpleCache[T any] struct {
	m    map[string]*SimpleCacheMap[T] // index 为 pageIndex + pageSize
	lock sync.RWMutex
	time time.Duration
}

type SimpleCacheMap[T any] struct {
	Data    *T
	EndTime int64
}

// NewSimpleCache
// @param time 缓存时间 time.Duration
func NewSimpleCache[T any](time time.Duration) *SimpleCache[T] {
	return &SimpleCache[T]{
		m:    make(map[string]*SimpleCacheMap[T]),
		time: time,
	}
}

// Get 获取
func (c *SimpleCache[T]) Get(key string) (*T, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	cache, ok := c.m[key]
	if ok && cache.EndTime >= time.Now().UnixNano() {
		return cache.Data, true
	}
	return nil, false
}

// Put 设置
func (c *SimpleCache[T]) Put(key string, t *T) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.m[key] = &SimpleCacheMap[T]{
		Data:    t,
		EndTime: time.Now().Add(c.time).UnixNano(),
	}
}

// PutLockFunc 设置并支持幂等性方法
func (c *SimpleCache[T]) PutLockFunc(key string, f func() (*T, error)) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	t, err := f()
	if err == nil {
		c.m[key] = &SimpleCacheMap[T]{
			Data:    t,
			EndTime: time.Now().Add(c.time).UnixNano(),
		}
	}
	return err
}

// IfNoExistPut 如果不存在则创建，存在则不创建，如redis setnx
// @return bool 是否存在
func (c *SimpleCache[T]) IfNoExistPut(key string, t *T) bool {
	_, ok := c.Get(key)
	if !ok {
		c.Put(key, t)
	}
	return ok
}

// GetOrCreate 存在则正常获取，不存在则创建默认值并返回
func (c *SimpleCache[T]) GetOrCreate(key string, t *T) *T {
	_ = c.IfNoExistPut(key, t)
	return t
}
