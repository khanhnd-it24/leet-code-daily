package design_hashmap

// using map
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

// using linked list
const size = 10001

type MyHashMap2 struct {
	list [size]*Node
}

type Node struct {
	key, val int
	next     *Node
}

func Constructor2() MyHashMap2 {
	return MyHashMap2{}
}

func (this *MyHashMap2) Put(key int, value int) {
	hashed := key % size
	if this.list[hashed] == nil {
		this.list[hashed] = &Node{key, value, nil}
		return
	}

	node := this.list[hashed]
	if node.key == key {
		node.val = value
		return
	}

	for node.next != nil {
		if node.next.key == key {
			node.next.val = value
			return
		}
		node = node.next
	}

	node.next = &Node{key, value, nil}
}

func (this *MyHashMap2) Get(key int) int {
	hashed := key % size
	for node := this.list[hashed]; node != nil; node = node.next {
		if node.key == key {
			return node.val
		}
	}

	return -1
}

func (this *MyHashMap2) Remove(key int) {
	hashed := key % size
	node := this.list[hashed]
	if node == nil {
		return
	}

	if node.key == key {
		this.list[hashed] = node.next
		return
	}

	for node.next != nil {
		if node.next.key == key {
			node.next = node.next.next
			return
		}
		node = node.next
	}
}
