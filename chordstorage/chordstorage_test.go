package chordstorage_test

import (
	"log"
	"sort"
	"testing"
	"triblab"
)

func TestChordStorage(t *testing.T) {
	// var backs = []string{"localhost:12461",
	// 	"localhost:12462",
	// 	"localhost:12463",
	// 	"localhost:12465",
	// 	"localhost:12466",
	// 	"localhost:12467"}

	var backs = []string{"localhost:1234", "localhost:5678", "localhost:123"}

	var backs2 = make([]string, len(backs))
	copy(backs2, backs)
	sort.Sort(triblab.ByHash{backs2})

	var storage = triblab.ChordStorage{Backs: backs2}

	for _, item := range backs2 {
		log.Println("item:", item, "hash:", storage.Hash(item))
	}

	addr := func(input string) string {
		return storage.FindMatchingNodeAddr(input)
	}

	if addr("h8liu") != backs[0] {
		t.Fail()
	}

	if addr("localhost:1234") != backs[0] {
		t.Fail()
	}

	if addr("rkapoor") != backs[1] {
		t.Fail()
	}

	if storage.FindSuccessor([]byte{245, 43, 103, 220, 74, 62, 71, 119, 93, 50, 189, 179,
		14, 169, 191, 43, 177, 101, 104, 32}) != backs[1] {
		t.Fail()
	}
}
