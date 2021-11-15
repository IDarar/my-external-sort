package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func Sort(f *os.File, fileSize int) {

	end := 0
	start := 0

	for i := 0; i < partition; i++ {
		end = (fileSize / partition) * (i + 1)

		nums := make([]int, 0)

		buf := make([]byte, end-start)

		n, err := f.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		start += n
		if buf[len(buf)-1] != '\n' {
			for {
				endLineBuf := make([]byte, 1)
				n, err := f.Read(endLineBuf)
				if err != nil && err != io.EOF {
					log.Fatal(err)
				}

				start += n

				if endLineBuf[len(endLineBuf)-1] == '\n' {
					buf = append(buf, endLineBuf...)
					break
				} else {
					buf = append(buf, endLineBuf...)
					continue
				}

			}
		}

		sp := bytes.Split(buf, []byte("\n"))
		for _, v := range sp {
			if len(v) > 0 {
				i, err := strconv.Atoi(string(v))
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, i)
			}
		}

		//sort nums
		quickSort(nums, 0, len(nums)-1)

		if err = writeNumsToFile(nums, "part"+fmt.Sprint(i)); err != nil {
			log.Fatal(err)
		}

	}

	readParts()
}
