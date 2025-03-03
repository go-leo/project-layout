package cachex

import (
	"sync"
	"time"

	"github.com/allegro/bigcache"
	"github.com/go-leo/cache"
	lru "github.com/hashicorp/golang-lru"
)

// Store 定义接口
type Store interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}

// sampleCache 简单缓存
type sampleCache struct {
	Cache sync.Map
}

func (store *sampleCache) Get(key string) (interface{}, bool) {
	return store.Cache.Load(key)
}

func (store *sampleCache) Set(key string, val interface{}) {
	store.Cache.Store(key, val)
}

// ttlCache TTL缓存
type ttlCache struct {
	Cache cache.Cache
	// 过期时间
	TTL func(key string) time.Duration
}

func (store *ttlCache) Get(key string) (interface{}, bool) {
	return store.Cache.Get(key)
}

func (store *ttlCache) Set(key string, val interface{}) {
	ttl := cache.DefaultExpiration
	if store.TTL != nil {
		ttl = store.TTL(key)
	}
	store.Cache.Set(key, val, ttl)
}

// lruCache LRU缓存
type lruCache struct {
	Cache *lru.Cache
}

func (store *lruCache) Get(key string) (interface{}, bool) {
	return store.Cache.Get(key)
}

func (store *lruCache) Set(key string, val interface{}) {
	store.Cache.Add(key, val)
}

// bigCache big cache
type bigCache struct {
	Cache      *bigcache.BigCache
	Marshal    func(key string, obj interface{}) ([]byte, error)
	Unmarshal  func(key string, data []byte) (interface{}, error)
	ErrHandler func(err error)
}

func (store *bigCache) Get(key string) (interface{}, bool) {
	data, err := store.Cache.Get(key)
	if err != nil {
		return nil, false
	}
	if store.Unmarshal == nil {
		return data, true
	}
	obj, err := store.Unmarshal(key, data)
	if err != nil {
		return nil, false
	}
	return obj, true
}

func (store *bigCache) Set(key string, val interface{}) {
	var err error
	switch v := val.(type) {
	case []byte:
		err = store.Cache.Set(key, v)
	case string:
		err = store.Cache.Set(key, []byte(v))
	default:
		if store.Unmarshal == nil {
			return
		}
		if data, e := store.Marshal(key, val); e != nil {
			err = e
		} else {
			err = store.Cache.Set(key, data)
		}
	}
	if err != nil && store.ErrHandler != nil {
		store.ErrHandler(err)
	}
}

func SampleCache() Store {
	return &sampleCache{Cache: sync.Map{}}
}

func TTLCache(
	expirationTime, cleanupInterval time.Duration,
	shards int,
	onEvicted func(string, interface{}),
	ttl func(key string) time.Duration,
) Store {
	c := cache.New(
		cache.ExpirationTime(expirationTime),
		cache.CleanupInterval(cleanupInterval),
		cache.Shards(shards),
		cache.OnEvicted(onEvicted),
	)
	return &ttlCache{
		Cache: c,
		TTL:   ttl,
	}
}

func LRUCache(
	size uint,
	onEvicted func(key interface{}, value interface{}),
) (Store, error) {
	evict, err := lru.NewWithEvict(int(size), onEvicted)
	if err != nil {
		return nil, err
	}
	return &lruCache{Cache: evict}, nil
}

func BigCache(
	expirationTime time.Duration,
	cleanupInterval time.Duration,
	shards int,
	onEvicted func(key string, entry []byte),
	marshal func(key string, obj interface{}) ([]byte, error),
	unmarshal func(key string, data []byte) (interface{}, error),
	errHandler func(err error),
) (Store, error) {
	config := bigcache.DefaultConfig(expirationTime)
	config.CleanWindow = cleanupInterval
	config.Shards = shards
	config.OnRemove = onEvicted
	c, err := bigcache.NewBigCache(config)
	if err != nil {
		return nil, err
	}
	return &bigCache{Cache: c, Marshal: marshal, Unmarshal: unmarshal, ErrHandler: errHandler}, nil
}
