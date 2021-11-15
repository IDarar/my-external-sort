package main

import (
	"os"
	"testing"
)

//Is faster then concurrent on hdd
func BenchmarkElemCounter(b *testing.B) {
	f, err := os.OpenFile("source", os.O_RDONLY, 0644)
	if err != nil {
		b.Fatal(err)
	}

	elemCounter(f)
}

//Is faster then sequential on ssd
func BenchmarkConcElemCounter(b *testing.B) {
	f, err := os.OpenFile("source", os.O_RDONLY, 0644)
	if err != nil {
		b.Fatal(err)
	}

	s, err := f.Stat()
	if err != nil {
		b.Fatal(err)
	}

	defer f.Close()

	concElemCounter(int(s.Size()), 4)
}
