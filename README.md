# Bloom Words

A lightweight Go library for efficient English word validation using [Bloom filters](https://en.wikipedia.org/wiki/Bloom_filter). Perfect for spell-checking, word games, and text validation with minimal memory footprint.

[![Go Reference](https://pkg.go.dev/badge/github.com/oosawy/bloomwords.svg)](https://pkg.go.dev/github.com/oosawy/bloomwords)

## What is Bloom Words?

Bloom Words is a Go library that validates English words using Bloom filters—achieving fast lookups with minimal memory usage. Perfect for spell-checking, word validation, and text filtering.

## Features

- 🚀 **Fast Lookup**: O(1) constant-time word lookup using Bloom filter
- 💾 **Memory Efficient**: Compressed filter using bitsets, much smaller than storing all words
- 📖 **Common English Words**: Pre-built filter with top 10,000 English words
- 📦 **Lightweight**: Entire filter embedded in binary, only ~12KB
- 🧪 **Well Tested**: Includes comprehensive test suite

**Quick Stats:**

- **10,000 common English words** compressed into **~12KB**
- **Sub-microsecond lookups** - test a word in less than 1 microsecond
- **Minimal false positive rate**: **~1%**, optimized for top common words
- **Zero false negatives** - if a word exists, you'll always find it

## Installation

```bash
go get github.com/oosawy/bloomwords
```

## Usage

### Basic Word Lookup

```go
package main

import (
	"fmt"
	"log"

	"github.com/oosawy/bloomwords"
)

func main() {
	// Initialize the Bloom filter
	bw, err := bloomwords.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Test if a word exists in the dictionary
	if bw.Test("hello") {
		fmt.Println("'hello' is a valid word")
	}

	if !bw.Test("xyzabc") {
		fmt.Println("'xyzabc' is likely not a valid word")
	}
}
```

## How It Works

Bloom Words uses Go's `go:embed` directive to embed the pre-built Bloom filter (`filter/bloom_words.bf`) directly into the binary. This eliminates the need to load external files at runtime and removes external dependencies. The embedded filter is loaded into memory during initialization, and all subsequent word lookups execute in constant O(1) time against this in-memory data.

## Building the Filter

To rebuild the Bloom filter from the word list:

```bash
go run ./cmd/build/build.go
```

This reads from `datasets/common_english_words.txt` and generates a new `filter/bloom_words.bf`.

### Dataset

The English word dataset used in this project is sourced from [Common English words on Kaggle](https://www.kaggle.com/datasets/vaskon/common-english-words).

## Testing

Run the test suite:

```bash
go test -v
```

## License

MIT
