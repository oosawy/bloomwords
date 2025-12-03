# Bloom Words

A lightweight Go library for efficient English word validation using [Bloom filters](https://en.wikipedia.org/wiki/Bloom_filter). Perfect for spell-checking, word games, and text validation with minimal memory footprint.

[![Go Reference](https://pkg.go.dev/badge/github.com/oosawy/bloom-words.svg)](https://pkg.go.dev/github.com/oosawy/bloom-words)

## What is Bloom Words?

Bloom Words is a Go library that validates English words using Bloom filters—achieving fast lookups with minimal memory usage. Perfect for spell-checking, word validation, and text filtering.

## Features

- 🚀 **Fast Lookup**: O(1) constant-time word lookup using Bloom filter
- 💾 **Memory Efficient**: Compressed filter using bitsets, much smaller than storing all words
- 📖 **Comprehensive Dictionary**: Pre-built filter with 370,000+ English words
- ⚡ **Streaming Support**: Efficiently handle large datasets with minimal memory usage
- 🧪 **Well Tested**: Includes comprehensive test suite

**Quick Stats:**

- **370K+ English words** compressed into just **~500KB** of data
- **Sub-microsecond lookups** - test a word in less than 1 microsecond
- **99% accuracy** - only 1% false positive rate on average
- **Zero false negatives** - if a word exists, you'll always find it

## Installation

```bash
go get github.com/oosawy/bloom-words
```

## Usage

### Basic Word Lookup

```go
package main

import (
	"fmt"
	"log"

	bw "github.com/oosawy/bloom-words"
)

func main() {
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
go run ./cmd/build.go
```

This reads from `datasets/words_alpha.txt` and generates a new `filter/bloom_words.bf`.

### Dataset

The English word dataset used in this project is sourced from [dwyl/english-words](https://github.com/dwyl/english-words).

## Testing

Run the test suite:

```bash
go test ./tests -v
```

## License

MIT
