package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func concElemCounter(size, concRate int) (int, error) {
	lineSep := []byte{'\n'}

	cChan := make(chan int, concRate)

	for i := 0; i < concRate; i++ {
		start := i * (size / concRate)
		end := (size / concRate) * (i + 1)

		go func(start, end int, i int) {
			f, err := os.OpenFile("source", os.O_RDONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}

			f.Seek(int64(start), 0)
			defer f.Close()
			buf := make([]byte, 1024*1024)

			gCount := 0

			for {
				//Check if goroitine is still in its bounds, not to count extra bytes
				if start+1024*1024 > end {
					c, err := f.Read(buf)

					//end - start is its part of file
					gCount += bytes.Count(buf[:end-start], lineSep)

					start += c

					switch {
					case err == io.EOF:

						cChan <- gCount

						return

					case err != nil:
						return
					}
				}
				if start >= end {
					cChan <- gCount
					return
				}
				c, err := f.ReadAt(buf, int64(start))
				start += c

				gCount += bytes.Count(buf[:c], lineSep)

				switch {
				case err == io.EOF:

					cChan <- gCount

					return

				case err != nil:
					return
				}
			}
		}(start, end, i)

	}

	count := 0

	rescOunt := 0
	for v := range cChan {
		count += v
		rescOunt++
		if rescOunt >= concRate {
			break
		}

	}

	return count, nil
}

func elemCounter(r io.Reader) (int, error) {

	lineSep := []byte{'\n'}
	count := 0

	buf := make([]byte, 1024*1024)

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
