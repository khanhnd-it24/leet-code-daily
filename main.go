package main

import "leet-code/src/design_hashmap"

func main() {
	hashMap := design_hashmap.Constructor2()

	hashMap.Put(1, 1)
	hashMap.Put(3, 4)
	hashMap.Put(2, 2)

	hashMap.Get(2)
	hashMap.Get(3)

	hashMap.Remove(2)
	hashMap.Get(2)
}
