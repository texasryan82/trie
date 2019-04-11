package trie

import "testing"

var (
	wordsAlphabet = []string{
		"alpha", "bravo", "charlie", "delta", "echo",
		"foxtrot", "golf", "hotel", "india", "juliett", "kilo", "lima",
		"mike", "november", "oscar", "papa", "quebec", "romeo", "sierra",
		"tango", "uniform", "victor", "whiskey", "xray", "yankee", "zulu",
	}
	wordsReverseAlphabet = []string{
		"zulu", "yankee", "xray", "whiskey", "victor", "uniform", "tango",
		"sierra", "romeo", "quebec", "papa", "oscar", "november", "mike",
		"lima", "kilo", "juliett", "india", "hotel", "golf", "foxtrot",
		"echo", "delta", "charlie", "bravo", "alpha"}
)

func TestCreate(t *testing.T) {
	trie := NewTrie()

	if trie == nil {
		t.Error("trie cannot be nil")
	}

	if trie.Count() != 0 {
		t.Error("trie should have zero words after creation")
	}
}

func TestInsertEmptyWord(t *testing.T) {
	trie := NewTrie()

	trie.Insert("")

	if trie.Count() != 0 {
		t.Error("trie should not insert empty word")
	}
}

func TestShouldNotContainEmptyWord(t *testing.T) {
	trie := NewTrie()

	if trie.Contains("") {
		t.Error("trie should not contain empty word")
	}
}

func TestShouldNotContainWordWhenEmpty(t *testing.T) {
	trie := NewTrie()

	if trie.Contains("whatever") {
		t.Error("trie should not contain \"whatever\"")
	}
}

func TestInsertWord(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foobar")
	if trie.Count() != 1 {
		t.Error("trie should have one word")
	}

	if !trie.Contains("foobar") {
		t.Error("trie should contain foobar")
	}

	if trie.Contains("foo") {
		t.Error("trie should not contain foo")
	}
}

func TestInsertOverlappingWords(t *testing.T) {
	trie := NewTrie()

	trie.Insert("foobar")
	trie.Insert("foo")

	if trie.Count() != 2 {
		t.Error("trie should have two words")
	}

	if !trie.Contains("foobar") {
		t.Error("trie should contain foobar")
	}

	if !trie.Contains("foo") {
		t.Error("trie should contain foo")
	}
}

func TestWordsInsertedInOrder(t *testing.T) {
	trie := NewTrie()

	for _, w := range wordsAlphabet {
		trie.Insert(w)
		if !trie.Contains(w) {
			t.Errorf("trie should contain %v", w)
		}
	}

	if trie.Count() != 26 {
		t.Error("trie should have 26 words")
	}
}

func TestWordsInsertedOutOfOrder(t *testing.T) {
	trie := NewTrie()

	for _, w := range wordsReverseAlphabet {
		trie.Insert(w)
		if !trie.Contains(w) {
			t.Errorf("trie should contain %v", w)
		}
	}

	if trie.Count() != 26 {
		t.Error("trie should have 26 words")
	}
}
