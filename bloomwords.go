package bloomwords

import (
	"bytes"

	"github.com/bits-and-blooms/bloom/v3"
)

const FilterFile = "filter/bloom_words.bf"

var filter bloom.BloomFilter

// BloomWords validates English words using a Bloom filter.
type BloomWords struct{}

// Init initializes BloomWords and returns a ready-to-use validator.
func Init() (*BloomWords, error) {
	r := bytes.NewReader(BloomWordsFilter)
	_, err := filter.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	BloomWordsFilter = nil // free memory
	return &BloomWords{}, nil
}

// Test returns true if the word is in the dictionary.
func (bw *BloomWords) Test(word string) bool {
	return filter.Test([]byte(word))
}
