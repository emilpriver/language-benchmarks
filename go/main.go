package main

import (
	"fmt"
	"sync"
)

func calculateSumOfSquares(chunk []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	sum := 0
	for _, x := range chunk {
		sum += x * x
	}
	ch <- sum
}

func main() {
	numbers := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		numbers[i] = i + 1
	}

	chunkSize := 10_000
	chunks := chunkList(numbers, chunkSize)

	var wg sync.WaitGroup
	ch := make(chan int, len(chunks))

	for _, chunk := range chunks {
		wg.Add(1)
		go calculateSumOfSquares(chunk, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalSum := 0
	for sum := range ch {
		totalSum += sum
	}

	fmt.Printf("Sum of squares: %d\n", totalSum)
}

func chunkList(list []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(list); i += chunkSize {
		end := i + chunkSize
		if end > len(list) {
			end = len(list)
		}
		chunks = append(chunks, list[i:end])
	}
	return chunks
}
