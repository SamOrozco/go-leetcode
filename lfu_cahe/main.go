package main

func main() {

}

type LFUCache struct {
	data     map[int]int
	capacity int
	size     int
	keyStack stack
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		data:     make(map[int]int, 0),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if val, ok := this.data[key]; ok {
		return val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	sz := this.size
	cp := this.capacity
	// if capacity get last used
	if sz == cp {
		_, val := this.keyStack.Pop()
		// remove entry from map
		delete(this.data, val)
		sz--
	}

}

func (this *LFUCache) evictLF() {

}
