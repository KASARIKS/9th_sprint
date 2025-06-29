package main

import (
	"math"
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	t.Parallel()
	t.Run("zero size", func(t *testing.T) {
		sliceLen := 0
		slice, err := generateRandomElements(sliceLen)

		if len(slice) != sliceLen {
			t.Errorf("Expected 0 length of slice but got %d\n", len(slice))
		}

		if err != nil {
			t.Errorf("Expected nil in error but got %s\n", err.Error())
		}
	})

	t.Run("minus size", func(t *testing.T) {
		slice, err := generateRandomElements(-10)

		if err == nil {
			t.Errorf("Expected error but got nil\n")
		}

		if len(slice) != 0 {
			t.Errorf("Expected 0 length of slice but got %d\n", len(slice))
		}
	})

	t.Run("normal size", func(t *testing.T) {
		sliceLen := 10
		slice, err := generateRandomElements(sliceLen)

		if err != nil {
			t.Errorf("Expected nil in error but got %s\n", err.Error())
		}

		if len(slice) != sliceLen {
			t.Errorf("Expected %d length of slice but got %d\n", sliceLen, len(slice))
		}

		// Didn't invent better way to check slice filling
		if slice[0] == slice[1] {
			t.Errorf("Expected first and second elements not equil. But got %d == %d\n."+
				"Could be a wrong result, restart test and check.\n", slice[0], slice[1])
		}
	})

	// Max size is int16
	t.Run("edge size", func(t *testing.T) {
		sliceLen := math.MaxInt16
		slice, err := generateRandomElements(sliceLen)

		if err != nil {
			t.Errorf("Expected nil in error but got %s\n", err.Error())
		}

		if len(slice) != sliceLen {
			t.Errorf("Expected %d length of slice but got %d\n", sliceLen, len(slice))
		}

		// Didn't invent better way to check slice filling
		if slice[0] == slice[1] {
			t.Errorf("Expected first and second elements not equil. But got %d == %d\n."+
				"Could be a wrong result, restart test and check.\n", slice[0], slice[1])
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("0 length", func(t *testing.T) {
		slice := []int{}
		maxNum := maximum(slice)

		if maxNum != 0 {
			t.Errorf("Expected 0 got %d\n", maxNum)
		}
	})

	t.Run("10 length", func(t *testing.T) {
		slice := []int{-5, 1, 2, 3, 10, -20, 10, 100, 1000, 200}
		maxNum := maximum(slice)

		if maxNum != 1000 {
			t.Errorf("Expected 1000 got %d", maxNum)
		}
	})

	t.Run("32767 length", func(t *testing.T) {
		slice := make([]int, math.MaxInt16)
		for i := range math.MaxInt16 {
			slice[i] = i
		}
		maxNum := maximum(slice)

		if maxNum != math.MaxInt16-1 {
			t.Errorf("Expected 1000 got %d", maxNum)
		}
	})
}
