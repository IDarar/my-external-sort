package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func quickSort(a []int, low, up int) {
	//Base case
	//If ephemeral array contains
	if up <= low {
		return
	}

	//take point from which we are going to devide initial slice
	//it can be any point in array
	//at the end of sorting it will take its correct position in whole array
	//left part of array is smaller, right is greater
	pivot := (low + up) / 2

	//start point represents position+1 of last number that is less than pivot
	start := low

	a[up], a[pivot] = a[pivot], a[up]

	for j := low; j < up; j++ {
		if a[j] < a[up] {
			a[start], a[j] = a[j], a[start]
			start++
		}
	}

	a[start], a[up] = a[up], a[start]

	quickSort(a, low, start-1)
	quickSort(a, start+1, up)
}

type merger struct {
	lastPopedQueueIndex int
	queues              []queue
	resuletFile         *os.File
}

type queue struct {
	f        *os.File
	isClosed bool
}

func (q *queue) pop() (int, bool) {
	if !q.isClosed {

		var num int
		resBuf := make([]byte, 0)
		for {
			buf := make([]byte, 1)
			//read file by one byte until we get whole number
			_, err := q.f.Read(buf)
			if err != nil {
				if err == io.EOF {
					q.isClosed = true
					return 0, false
				}
				log.Fatal(err)
			}

			if buf[len(buf)-1] == '\n' {
				number, err := strconv.Atoi(string(resBuf))
				if err != nil {
					log.Fatal(err)
				}

				num = number
				break
			} else {
				resBuf = append(resBuf, buf...)
				continue
			}
		}

		return num, true
	}

	return 0, false
}

func (m merger) getValueFromNextAvailableQueue() (minHeapNode, bool) {
	for i := range m.queues {
		if !m.queues[i].isClosed {
			num, ok := m.queues[i].pop()
			if !ok {
				continue
			}
			return minHeapNode{value: num, queue: i}, true
		}
	}
	return minHeapNode{}, false
}

func readParts() {
	m := merger{}

	mHeap := []minHeapNode{}

	res, err := os.Create("result")
	if err != nil {
		log.Fatal(err)
	}

	m.resuletFile = res

	for i := 0; i < partition; i++ {
		q := queue{}

		f, err := os.Open("part" + fmt.Sprint(i))
		if err != nil {
			log.Fatal(err)
		}
		q.f = f

		m.queues = append(m.queues, q)

		num, ok := q.pop()
		if !ok {
			continue
		}

		val := minHeapNode{queue: i, value: num}

		mHeap = insert(minHeapNode{value: num, queue: val.queue}, mHeap)
	}

	leftQueues := partition

	counter := -1
	var val minHeapNode

	for leftQueues > 0 {
		counter++
		val, mHeap = removeRoot(mHeap)

		m.writeToRes(val.value)

		num, ok := m.queues[val.queue].pop()
		if !ok {
			leftQueues--
			if leftQueues > 0 {
				val, ok = m.getValueFromNextAvailableQueue()
				if !ok {
					break
				}
				num = val.value
			} else {
				break
			}

		}

		mHeap = insert(minHeapNode{value: num, queue: val.queue}, mHeap)
	}

	//write values left in heap
	left := len(mHeap)
	for i := 0; i < left; i++ {
		val, mHeap = removeRoot(mHeap)
		m.writeToRes(val.value)
	}
}

func (m merger) writeToRes(num int) {
	_, err := m.resuletFile.Write([]byte(fmt.Sprintln(num)))
	if err != nil {
		log.Fatal(err)
	}
}
