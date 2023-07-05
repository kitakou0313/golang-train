package main

func main() {
	lfu := &Lfu{}
	lru := &Lru{}
	fifo := &Fifo{}

	cache := initCache(lfu)

	cache.add("A", "1")
	cache.add("B", "2")

	cache.add("C", "3")

	cache.setEvictionAlgo(lru)
	cache.add("D", "4")

	cache.setEvictionAlgo(fifo)
	cache.add("E", "5")
}
