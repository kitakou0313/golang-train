package main

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	mapCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)

	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		mapCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.mapCapacity {
		c.evict()
	}

	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}
