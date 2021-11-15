package main

type minHeapNode struct {
	value int
	queue int
}

func downHeapify(a []minHeapNode, current int) {
	if current >= (len(a)/2) && current <= len(a) {
		return
	}

	smallest := current
	leftChildIndex := 2*current + 1
	rightRightIndex := 2*current + 2

	if leftChildIndex < len(a) && a[leftChildIndex].value < a[smallest].value {
		smallest = leftChildIndex
	}
	if rightRightIndex < len(a) && a[rightRightIndex].value < a[smallest].value {
		smallest = rightRightIndex
	}
	if smallest != current {

		a[current], a[smallest] = a[smallest], a[current]

		downHeapify(a, smallest)
	}

}
func heapifyFromBottom(arr []minHeapNode, index int) []minHeapNode {
	parent := (index - 1) / 2

	if parent < 0 {
		return arr
	}

	if arr[index].value < arr[parent].value {
		arr[index], arr[parent] = arr[parent], arr[index]

		heapifyFromBottom(arr, parent)
	}
	return arr
}
func removeRoot(arr []minHeapNode) (minHeapNode, []minHeapNode) {
	root := arr[0]

	arr[0] = arr[len(arr)-1]

	arr = arr[:len(arr)-1]

	downHeapify(arr, 0)

	return root, arr
}

func insert(val minHeapNode, arr []minHeapNode) []minHeapNode {
	arr = append(arr, val)

	heapifyFromBottom(arr, len(arr)-1)

	return arr
}
