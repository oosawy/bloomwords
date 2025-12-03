package bloomwords

import _ "embed"

//go:embed filter/bloom_words.bf
var BloomWordsFilter []byte
