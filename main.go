package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
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
func maxChunks(data []int) (int, error) {
	sliceLen := len(data)

	if sliceLen < CHUNKS {
		return maximum(data), nil
	}

	var wg sync.WaitGroup

	dividedMaximums := make([]int, CHUNKS)

	var i int

	for i = 0; i < CHUNKS-1; i++ {
		idx1 := sliceLen / CHUNKS * i
		idx2 := sliceLen / CHUNKS * (i + 1)
		slicePart := data[idx1:idx2]

		wg.Add(1)
		go func(slicePart []int, i int) {
			defer wg.Done()
			dividedMaximums[i] = maximum(slicePart)
		}(slicePart, i)
	}

	wg.Wait()

	idx1 := sliceLen / CHUNKS * (i + 1)
	dividedMaximums[i] = maximum(data[idx1:])

	return maximum(dividedMaximums), nil
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
	max, _ = maxChunks(numbers)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
