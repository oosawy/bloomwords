package main

import (
	"bufio"
	"iter"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bits-and-blooms/bloom/v3"
)

const DatasetWordsAlpha = "datasets/words_alpha.txt"
const DatasetWordsAlphaCount = "datasets/words_alpha_count"
const OutputBloomFilter = "filter/bloom_words.bf"

func main() {
	count, seq := load_dataset()
	log.Printf("Total words: %d\n", count)

	filter := bloom.NewWithEstimates(uint(count), 0.01)

	added := 0
	for word := range seq {
		added++
		filter.Add(word)

		if added%10000 == 0 {
			percent := float64(added) / float64(count) * 100
			log.Printf("Added %d words (%.2f%%)", added, percent)
		}
	}
	log.Printf("Finished adding words. Total added: %d\n", added)

	out, err := os.Create(OutputBloomFilter)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	wrote, err := filter.WriteTo(writer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes\n", wrote)

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bloom filter saved to bloom_words.bf")
}

func load_dataset() (count int, seq iter.Seq[[]byte]) {
	f, err := os.ReadFile(DatasetWordsAlphaCount)
	if err != nil {
		log.Fatal(err)
	}

	count, err = strconv.Atoi(strings.TrimSpace(string(f)))
	if err != nil {
		log.Fatal(err)
	}

	seq = func(yield func([]byte) bool) {
		f, err := os.Open(DatasetWordsAlpha)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if !yield(scanner.Bytes()) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return count, seq
}
