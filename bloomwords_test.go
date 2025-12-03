package bloomwords_test

import (
	"testing"

	bloomwords "github.com/oosawy/bloom-words"
)

func TestBloomWordsWithCommonWords(t *testing.T) {
	bw, err := bloomwords.Init()
	if err != nil {
		t.Fatalf("Failed to initialize BloomWords: %v", err)
	}

	commonWords := []string{
		"hello",
		"world",
		"computer",
		"programming",
		"test",
		"function",
		"variable",
		"algorithm",
		"database",
		"network",
	}

	for _, word := range commonWords {
		if !bw.Test(word) {
			t.Errorf("Expected word '%s' to be found in bloom filter, but it was not", word)
		}
	}
}

func TestBloomWordsWithNonExistentWords(t *testing.T) {
	bw, err := bloomwords.Init()
	if err != nil {
		t.Fatalf("Failed to initialize BloomWords: %v", err)
	}

	nonExistentWords := []string{
		"xyzabc",
		"qwertyu",
		"asdfghj",
		"zxcvbnm",
		"abcdefghijklmnopqrstuvwxyz",
		"aaaaaaaaa",
		"zzzzzzzzz",
		"notarealword",
	}

	for _, word := range nonExistentWords {
		if bw.Test(word) {
			t.Logf("Warning: non-existent word '%s' was found (false positive)", word)
		}
	}
}
