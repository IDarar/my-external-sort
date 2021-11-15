package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

const (
	elemNumber = 500000
)

//assume that file to be sorted is just text integers separated by newline
func genFile() {
	//set up test file
	f, err := os.OpenFile("source", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < elemNumber; i++ {
		val := rand.Intn(1000000000)

		if _, err := f.Write([]byte(fmt.Sprintln(val))); err != nil {
			log.Fatal(err)
		}
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

const partition = 4

func main() {
	f, err := os.OpenFile("source", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	Sort(f, int(s.Size()))
}
