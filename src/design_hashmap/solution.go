package design_hashmap

type MyHashMap struct {
	mp map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{mp: make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.mp[key] = value
}

func (this *MyHashMap) Get(key int) int {
	val, ok := this.mp[key]
	if !ok {
		return -1
	}
	return val
}

func (this *MyHashMap) Remove(key int) {
	delete(this.mp, key)
}
