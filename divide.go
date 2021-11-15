package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func isCorrectlyDivided() bool {
	f, err := os.OpenFile("source", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	partsSize := 0

	for i := 0; i < partition; i++ {
		partF, err := os.OpenFile("part"+fmt.Sprint(i), os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		s, err := partF.Stat()
		if err != nil {
			log.Fatal(err)
		}

		partBuf := make([]byte, s.Size())

		partsSize += int(s.Size())
		_, err = partF.Read(partBuf)
		if err != nil {
			log.Fatal(err)
		}
		fBuf := make([]byte, s.Size())
		_, err = f.Read(fBuf)
		if err != nil {
			log.Fatal(err)
		}

		eq := bytes.Equal(fBuf, partBuf)
		if !eq {
			return eq
		}
	}

	fileStat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if int(fileStat.Size()) != partsSize {
		fmt.Println("not eq")
	}

	return true
}

func writeNumsToFile(nums []int, filename string) error {
	partF, err := os.Create(filename)
	if err != nil {
		return err
	}

	//write sorted numbers into temp part file
	for _, v := range nums {
		if _, err := partF.Write([]byte(fmt.Sprintln(v))); err != nil {
			return err
		}
	}

	return nil
}
