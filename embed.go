package bloomwords

import _ "embed"

// BloomWordsFilter is a byte slice containing the serialized Bloom filter.
// It will be free'd after initialization to save memory.
//
//go:embed filter/bloom_words.bf
var BloomWordsFilter []byte
