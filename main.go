package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	//SIZE   = 100_000_000 True size
	SIZE   = 1000 // For single goroutine
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
	// ваш код здесь
	if size < 0 {
		return []int{}, fmt.Errorf("size smaller than 0")
	}

	generator := rand.New(rand.NewSource(time.Now().Unix()))
	randomElements := make([]int, size)
	for i := 0; i < size; i++ {
		randomElements[i] = generator.Int()
	}

	return randomElements, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	var max int
	for _, n := range data {
		if n > max {
			max = n
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь

	return data[0]
}

func main() {
	max, elapsed := 0, 0

	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	numbers, _ := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	maxNum := maximum(numbers)
	fmt.Println(maxNum)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	maxChunks(numbers)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
