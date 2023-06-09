package mappkg

import "sync"

type syncMap[K any, T any] struct {
	m sync.Map
}

func NewSyncMap[K any, T any]() *syncMap[K, T] {
	return &syncMap[K, T]{
		m: sync.Map{},
	}
}

func (s *syncMap[K, T]) Store(key K, value T) {
	s.m.Store(key, value)
}

func (s *syncMap[K, T]) Delete(key K) {
	s.m.Delete(key)
}

func (s *syncMap[K, T]) Load(key K) (T, bool) {
	load, ok := s.m.Load(key)
	if !ok || load == nil {
		var t T
		return t, ok
	}
	return load.(T), ok
}

func (s *syncMap[K, T]) LoadOrStore(key K, value T) (T, bool) {
	store, ok := s.m.LoadOrStore(key, value)
	if !ok || store == nil {
		var t T
		return t, ok
	}
	return store.(T), ok
}

func (s *syncMap[K, T]) LoadAndDelete(key K) (T, bool) {
	value, ok := s.m.LoadAndDelete(key)
	if !ok || value == nil {
		var t T
		return t, ok
	}
	return value.(T), ok
}

func (s *syncMap[K, T]) Range(f func(key K, value T) bool) {
	s.m.Range(func(key, value any) bool {
		if _, ok := key.(K); !ok {
			return true
		}
		if _, ok := value.(T); !ok {
			return true
		}

		if value == nil {
			var t T
			return f(key.(K), t)
		}
		return f(key.(K), value.(T))
	})
}
