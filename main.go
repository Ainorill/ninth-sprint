package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE           = 10_000_000
	CHUNKS         = 8
	maxRandomValue = 1_000_000_000
)

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxValue := maximum(numbers)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxValue = maxChunks(numbers)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, elapsed)
}

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = rand.Intn(maxRandomValue) + 1
	}
	return numbers
}

func maximum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxValue := numbers[0]
	for _, number := range numbers[1:] {
		if number > maxValue {
			maxValue = number
		}
	}
	return maxValue
}

func maxChunks(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	chunkSize := len(numbers) / CHUNKS
	chunkMaxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		if i == CHUNKS-1 {
			endIndex = len(numbers)
		}

		go func(chunkIndex int, chunk []int) {
			defer wg.Done()
			chunkMaxes[chunkIndex] = maximum(chunk)
		}(i, numbers[startIndex:endIndex])
	}

	wg.Wait()

	return maximum(chunkMaxes)
}
