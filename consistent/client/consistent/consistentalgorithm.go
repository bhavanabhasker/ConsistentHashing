package consistent

import (
	"errors"
	"hash/fnv"
	"sort"
	"strconv"
	"sync"
)

// implementation of consistent hashing alogorithm
//Replication factor is ignored
// Consistent holds the information about the members of the consistent hash circle.
type Consistent struct {
	Circle       map[uint32]string
	SortedHashes uints
	// for locki
	sync.RWMutex
}
type uints []uint32

var Err = errors.New("empty circle")

func CreateConsistent() *Consistent {
	c := new(Consistent)
	c.Circle = make(map[uint32]string)
	return c
}

func (c *Consistent) AddNode(node string) {
	c.Lock()
	defer c.Unlock()
	c.Circle[hash(getstringkey(node, 1))] = node
	c.sort()
}

func (c *Consistent) AddData(value string) (string, error) {
	c.RLock()
	defer c.RUnlock()
	if len(c.Circle) == 0 {
		return "", Err
	}
	key := hash(value)
	i := c.search(key)
	return c.Circle[c.SortedHashes[i]], nil
}
func (c *Consistent) search(key uint32) (i int) {
	// attach to the nearest node
	i = sort.Search(len(c.SortedHashes), func(x int) bool {
		return c.SortedHashes[x] > key
	})
	if i >= len(c.SortedHashes) {
		i = 0
	}
	return
}

func hash(data string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(data))
	return h.Sum32()
}
func getstringkey(data string, i int) string {
	return data + strconv.Itoa(i)
}
func (c *Consistent) sort() {
	hashes := c.SortedHashes[:0]

	for node := range c.Circle {
		hashes = append(hashes, node)
	}
	sort.Sort(hashes)
	c.SortedHashes = hashes
}
func (slice uints) Len() int {
	return len(slice)
}

func (slice uints) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice uints) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
func (c *Consistent) RemoveNode(node string) {
	delete(c.Circle, hash(getstringkey(node, 1)))
	c.sort()
}
