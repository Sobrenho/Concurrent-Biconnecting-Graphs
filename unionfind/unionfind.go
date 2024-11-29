package unionfind

import (
	"sync"
)

func keysOf[K comparable, T any](m map[K]T) []K {
	keys := make([]K, len(m))
	keyCount := 0
	for key := range m {
		keys[keyCount] = key
		keyCount++
	}
	return keys
}

type UnionFind struct {
	size             int
	parent           []int
	representedCount []int
	vertexLocks      []sync.Mutex
	representatives  map[int]bool
	lock             sync.Mutex
}

func NewUnionFind(size int) *UnionFind {

	unionFind := new(UnionFind)

	unionFind.size             = size
	unionFind.parent           = make([]int, size)
	unionFind.representedCount = make([]int, size) 
	unionFind.vertexLocks      = make([]sync.Mutex, size)
	unionFind.representatives  = make(map[int]bool, size)

	for i := int(0); i < size; i++ {
		unionFind.parent[i] = i
		unionFind.representedCount[i] = 1
		unionFind.representatives[i] = true
	}

	return unionFind
}

func (unionFind *UnionFind) Find(vertex int) int {

	unionFind.vertexLocks[vertex].Lock()
	defer unionFind.vertexLocks[vertex].Unlock()

	if unionFind.parent[vertex] == vertex {
		return vertex
	}

	unionFind.parent[vertex] = unionFind.Find(unionFind.parent[vertex])

	return unionFind.parent[vertex]
}

func (unionFind *UnionFind) Join(vertexA int, vertexB int) {

	representativeA := unionFind.Find(vertexA)
	representativeB := unionFind.Find(vertexB)

	if representativeA != representativeB {

		unionFind.vertexLocks[representativeA].Lock()
		unionFind.vertexLocks[representativeB].Lock()
		unionFind.lock.Lock()

		if unionFind.representedCount[representativeA] >= unionFind.representedCount[representativeB] {

			unionFind.representedCount[representativeA] += unionFind.representedCount[representativeB]
			unionFind.parent[representativeB] = representativeA
			delete(unionFind.representatives, representativeB)

		} else {

			unionFind.representedCount[representativeB] += unionFind.representedCount[representativeA]
			unionFind.parent[representativeA] = representativeB
			delete(unionFind.representatives, representativeA)

		}

		unionFind.lock.Unlock()
		unionFind.vertexLocks[representativeA].Unlock()
		unionFind.vertexLocks[representativeB].Unlock()

	}
}

func (unionFind *UnionFind) Representatives() []int {
	unionFind.lock.Lock()
	defer unionFind.lock.Unlock()
	return keysOf(unionFind.representatives)
}

func (unionFind *UnionFind) Size() int {
	return unionFind.size
}