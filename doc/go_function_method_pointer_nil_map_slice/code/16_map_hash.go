// A cyclic redundancy check (CRC) is an error-detecting code commonly used in digital
// networks and storage devices to detect accidental changes to raw data.
//
// https://github.com/golang/go/blob/master/src/hash/crc32/crc32.go
// https://github.com/golang/go/blob/master/src/runtime/hashmap.go

// A map is just a hash table.  The data is arranged
// into an array of buckets.  Each bucket contains up to
// 8 key/value pairs.  The low-order bits of the hash are
// used to select a bucket.  Each bucket contains a few
// high-order bits of each hash to distinguish the entries
// within a single bucket.
//
// If more than 8 keys hash to a bucket, we chain on
// extra buckets.
//
// When the hashtable grows, we allocate a new array
// of buckets twice as big.  Buckets are incrementally
// copied from the old bucket array to the new bucket array.

package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"strings"
)

func main() {
	// hash table using array
	ht := newHashTable()
	for _, elem := range strings.Split("aaaaaaaaaaabbbbcdeftghiklmnopr", "") {
		ht.insert([]byte(elem))
	}
	for _, bucket := range ht.bucketSlice {
		fmt.Println(bucket)
	}
	/*
	   &{true [[102] [98] [101] [97] [100] [116] [103] [99]]}
	   &{true [[109] [105] [110] [112] [111] [107] [108] [104]]}
	   &{false [[] [] [] [] [] [114] [] []]}
	*/

	fmt.Println(ht.search([]byte("f"))) // true
	fmt.Println(ht.search([]byte("x"))) // false
}

func hashFuncCrc32(val []byte) uint32 {
	// crc64.Checksum(val, crc64.MakeTable(crc64.ISO))
	return crc32.Checksum(val, crc32.MakeTable(crc32.IEEE))
}

func hashFunc(val []byte) uint32 {
	return checksum(val, makePolyTable(crc32.IEEE))
}

// polyTable is a 256-word table representing the polynomial for efficient processing.
type polyTable [256]uint32

func makePolyTable(poly uint32) *polyTable {
	t := new(polyTable)
	for i := 0; i < 256; i++ {
		crc := uint32(i)
		for j := 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		t[i] = crc
	}
	return t
}

// checksum returns the CRC-32 checksum of data
// using the polynomial represented by the polyTable.
func checksum(data []byte, tab *polyTable) uint32 {
	crc := ^uint32(0)
	for _, v := range data {
		crc = tab[byte(crc)^v] ^ (crc >> 8)
	}
	return ^crc
}

const (
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits // 8, Maximum number of key/value pairs a bucket can hold
)

type hashTable struct {
	bucketSlice []*bucket
}

func newHashTable() *hashTable {
	table := new(hashTable)
	// table.bucketSlice = make([]*bucket, hashTableSize)
	table.bucketSlice = []*bucket{}
	return table
}

type bucket struct {
	wrapped bool // already wrapped around from end of bucket array to beginning
	data    [bucketCnt][]byte
	// type byteData []byte
	// []byte == []uint8
}

func newBucket() *bucket {
	newBucket := &bucket{}
	newBucket.wrapped = false
	newBucket.data = [bucketCnt][]byte{}
	return newBucket
}

func (h *hashTable) search(val []byte) bool {
	if len(h.bucketSlice) == 0 {
		return false
	}
	probeIdx := hashFunc(val) % uint32(bucketCnt)
	for _, bucket := range h.bucketSlice {
		// check the probeIdx
		if bucket.data[probeIdx] != nil {
			if bytes.Equal(bucket.data[probeIdx], val) {
				return true
			}
		}
		// linear probe
		for idx, elem := range bucket.data {
			if uint32(idx) == probeIdx {
				continue
			}
			if bytes.Equal(elem, val) {
				return true
			}
		}
	}
	return false
}

// hashFunc -> probeIdx -> linear probe to fill up bucket
func (h *hashTable) insert(val []byte) {
	if h.search(val) {
		return
	}
	if len(h.bucketSlice) == 0 {
		h.bucketSlice = append(h.bucketSlice, newBucket())
	}
	probeIdx := hashFunc(val) % uint32(bucketCnt)
	isInserted := false
Loop:
	for _, bucket := range h.bucketSlice {
		// if the bucket is already full, skip it
		if bucket.wrapped {
			continue
		}
		// if the index is not taken yet, map it
		if bucket.data[probeIdx] == nil {
			bucket.data[probeIdx] = val
			isInserted = true
			break
		}
		// linear probe
		for idx, elem := range bucket.data {
			if uint32(idx) == probeIdx {
				continue
			}
			if elem == nil {
				bucket.data[idx] = val
				isInserted = true
				break Loop
			}
		}
		bucket.wrapped = true
	}
	if !isInserted {
		nb := newBucket()
		nb.data[probeIdx] = val
		h.bucketSlice = append(h.bucketSlice, nb)
	}
}
