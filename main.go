package main

import (
	"fmt"
	"math/rand"
	"sync"
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
	var wg sync.WaitGroup
	divided, _ := divideSlice(data, 8)
	maxGos := make([]int, CHUNKS)

	for i := 0; i < len(divided); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			maxGos = append(maxGos, maximum(divided[i]))
		}()
	}

	wg.Wait()
	return maximum(maxGos)
}

func divideSlice(slice []int, parts int) ([][]int, error) {
	if len(slice) < parts {
		return [][]int{}, fmt.Errorf("Slice lenght bigger than parts count")
	}

	var divided [][]int = make([][]int, parts)
	basePartLength := int(len(slice) / parts)
	i := 0

	for part := range parts {
		divided[part] = make([]int, basePartLength)

		// Don't add wait to this goroutine, because wg.add it'll broken.
		// As I guess, copy have it owns mutexes, or something.
		go copy(divided[part], slice[i:i+basePartLength])

		i += basePartLength
	}

	for i2 := i; i2 < len(slice); i2++ {
		divided[i2-i] = append(divided[i2-i], slice[i2])
	}

	return divided, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers, _ := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
