package bloomwords

import (
	"bytes"

	"github.com/bits-and-blooms/bloom/v3"
)

const FilterFile = "filter/bloom_words.bf"

type BloomWords struct {
	filter bloom.BloomFilter
}

func Init() (*BloomWords, error) {
	var filter bloom.BloomFilter

	r := bytes.NewReader(BloomWordsFilter)
	_, err := filter.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	return &BloomWords{filter: filter}, nil
}

func (bw *BloomWords) Test(word string) bool {
	return bw.filter.Test([]byte(word))
}
