package tree

import (
	"fmt"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()

	collections := []string{
		"go", "google", "gobuild", "gostart", "gorun", "gotest",
		"rust", "python", "happy",
	}

	for _, s := range collections {
		trie.Insert(s)
	}
	fmt.Println(trie.StartWith("xxx"))
}
