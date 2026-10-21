package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "нулевой размер", size: 0, want: 0},
		{name: "отрицательный размер", size: -5, want: 0},
		{name: "один элемент", size: 1, want: 1},
		{name: "обычный размер", size: 1000, want: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if got == nil {
				t.Fatalf("generateRandomElements(%d) вернула nil", tt.size)
			}
			if len(got) != tt.want {
				t.Errorf("generateRandomElements(%d): длина = %d, ожидалось %d", tt.size, len(got), tt.want)
			}
			for i, number := range got {
				if number <= 0 {
					t.Errorf("generateRandomElements(%d): элемент %d = %d, ожидалось положительное число", tt.size, i, number)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{name: "nil слайс", numbers: nil, want: 0},
		{name: "пустой слайс", numbers: []int{}, want: 0},
		{name: "один элемент", numbers: []int{7}, want: 7},
		{name: "максимум в начале", numbers: []int{9, 2, 3}, want: 9},
		{name: "максимум в середине", numbers: []int{1, 10, 3}, want: 10},
		{name: "максимум в конце", numbers: []int{1, 2, 11}, want: 11},
		{name: "одинаковые элементы", numbers: []int{5, 5, 5, 5}, want: 5},
		{name: "несколько максимумов", numbers: []int{4, 8, 2, 8}, want: 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.numbers); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.numbers, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{name: "nil слайс", numbers: nil, want: 0},
		{name: "пустой слайс", numbers: []int{}, want: 0},
		{name: "один элемент", numbers: []int{3}, want: 3},
		{name: "элементов меньше, чем частей", numbers: []int{1, 5, 2}, want: 5},
		{name: "длина равна числу частей", numbers: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "длина не делится на число частей, максимум в конце", numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, want: 11},
		{name: "максимум в начале", numbers: []int{50, 1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.numbers); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.numbers, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{1, 7, 8, 9, 100, 1001, 100_000} {
		numbers := generateRandomElements(size)
		if got, want := maxChunks(numbers), maximum(numbers); got != want {
			t.Errorf("размер %d: maxChunks = %d, maximum = %d", size, got, want)
		}
	}
}
