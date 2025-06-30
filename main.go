package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	SIZE = 100_000_000 // True size
	//SIZE   = 1000 // For single goroutine
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
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
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, n := range data[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.

// Rarely tests fail on negatives numbers, give 0 and I don't understand why.
func maxChunks(data []int) (int, error) {
	if len(data) < CHUNKS {
		return maximum(data), nil
	}

	dividedMaximums, err := dividedMaximums(data, CHUNKS)

	return maximum(dividedMaximums), err
}

func dividedMaximums(slice []int, parts int) ([]int, error) {
	if len(slice) < parts {
		return []int{}, fmt.Errorf("Slice length tbigger than parts count")
	}

	var dividedMaximums []int = make([]int, parts)
	sliceLen := len(slice)
	basePartLength := int(sliceLen / parts)
	i := 0

	for part := range parts - 1 {
		dividedMaximums[part] = maximum(slice[i : i+basePartLength])

		i += basePartLength
	}

	dividedMaximums[parts-1] = maximum(slice[i:sliceLen])

	return dividedMaximums, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers, err := generateRandomElements(SIZE)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max, err = maxChunks(numbers)
	if err != nil {
		log.Fatal(err)
	}

	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
