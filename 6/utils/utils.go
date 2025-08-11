package utils

import (
	"math/rand"
	"time"
)

// InitRandom инициализирует генератор случайных чисел
func InitRandom() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt генерирует случайное целое число в диапазоне [min, max]
func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

// RandomFloat генерирует случайное дробное число в диапазоне [min, max)
func RandomFloat(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	return min + rand.Float64()*(max-min)
}

// RandomSlice генерирует слайс из n случайных чисел в диапазоне [min, max]
func RandomSlice(n, min, max int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = RandomInt(min, max)
	}
	return slice
}

// RandomBool генерирует случайное булево значение
func RandomBool() bool {
	return rand.Intn(2) == 1
}

// RandomChoice выбирает случайный элемент из слайса строк
func RandomChoice(choices []string) string {
	if len(choices) == 0 {
		return ""
	}
	return choices[rand.Intn(len(choices))]
}