package utils

import (
	"testing"
)

func TestInitRandom(t *testing.T) {
	// Просто проверяем, что функция не вызывает панику
	InitRandom()
}

func TestRandomInt(t *testing.T) {
	InitRandom()
	
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"Обычный диапазон", 1, 10},
		{"Отрицательные числа", -10, -1},
		{"Смешанный диапазон", -5, 5},
		{"Одинаковые значения", 5, 5},
		{"Перевёрнутый диапазон", 10, 1},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				result := RandomInt(tt.min, tt.max)
				expectedMin := tt.min
				expectedMax := tt.max
				if tt.min > tt.max {
					expectedMin, expectedMax = tt.max, tt.min
				}
				if result < expectedMin || result > expectedMax {
					t.Errorf("RandomInt(%d, %d) = %d; вне диапазона [%d, %d]",
						tt.min, tt.max, result, expectedMin, expectedMax)
				}
			}
		})
	}
}

func TestRandomFloat(t *testing.T) {
	InitRandom()
	
	tests := []struct {
		name string
		min  float64
		max  float64
	}{
		{"Обычный диапазон", 0.0, 1.0},
		{"Большой диапазон", 0.0, 100.0},
		{"Отрицательные числа", -10.0, -1.0},
		{"Перевёрнутый диапазон", 10.0, 1.0},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				result := RandomFloat(tt.min, tt.max)
				expectedMin := tt.min
				expectedMax := tt.max
				if tt.min > tt.max {
					expectedMin, expectedMax = tt.max, tt.min
				}
				if result < expectedMin || result >= expectedMax {
					t.Errorf("RandomFloat(%f, %f) = %f; вне диапазона [%f, %f)",
						tt.min, tt.max, result, expectedMin, expectedMax)
				}
			}
		})
	}
}

func TestRandomSlice(t *testing.T) {
	InitRandom()
	
	tests := []struct {
		name string
		n    int
		min  int
		max  int
	}{
		{"Маленький слайс", 5, 1, 10},
		{"Большой слайс", 100, 0, 50},
		{"Пустой слайс", 0, 1, 10},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomSlice(tt.n, tt.min, tt.max)
			if len(result) != tt.n {
				t.Errorf("RandomSlice(%d, %d, %d): длина = %d; ожидалось %d",
					tt.n, tt.min, tt.max, len(result), tt.n)
			}
			
			for i, val := range result {
				if val < tt.min || val > tt.max {
					t.Errorf("RandomSlice: элемент [%d] = %d вне диапазона [%d, %d]",
						i, val, tt.min, tt.max)
				}
			}
		})
	}
}

func TestRandomBool(t *testing.T) {
	InitRandom()
	
	trueCount := 0
	falseCount := 0
	iterations := 1000
	
	for i := 0; i < iterations; i++ {
		if RandomBool() {
			trueCount++
		} else {
			falseCount++
		}
	}
	
	// Проверяем, что оба значения встречаются
	if trueCount == 0 {
		t.Error("RandomBool: никогда не возвращает true")
	}
	if falseCount == 0 {
		t.Error("RandomBool: никогда не возвращает false")
	}
	
	// Проверяем примерное распределение (с большим допуском)
	ratio := float64(trueCount) / float64(iterations)
	if ratio < 0.3 || ratio > 0.7 {
		t.Logf("Предупреждение: RandomBool распределение может быть смещено: true=%.2f%%", ratio*100)
	}
}

func TestRandomChoice(t *testing.T) {
	InitRandom()
	
	tests := []struct {
		name    string
		choices []string
		wantEmpty bool
	}{
		{"Обычный список", []string{"a", "b", "c", "d"}, false},
		{"Один элемент", []string{"single"}, false},
		{"Пустой список", []string{}, true},
		{"nil список", nil, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomChoice(tt.choices)
			
			if tt.wantEmpty && result != "" {
				t.Errorf("RandomChoice(%v) = %q; ожидалась пустая строка", tt.choices, result)
			}
			
			if !tt.wantEmpty && result == "" {
				t.Errorf("RandomChoice(%v) = пустая строка; ожидался элемент из списка", tt.choices)
			}
			
			if !tt.wantEmpty {
				found := false
				for _, choice := range tt.choices {
					if result == choice {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("RandomChoice(%v) = %q; не найден в исходном списке", tt.choices, result)
				}
			}
		})
	}
	
	// Тест на распределение
	choices := []string{"a", "b", "c"}
	counts := make(map[string]int)
	iterations := 300
	
	for i := 0; i < iterations; i++ {
		result := RandomChoice(choices)
		counts[result]++
	}
	
	// Проверяем, что все варианты встречаются
	for _, choice := range choices {
		if counts[choice] == 0 {
			t.Errorf("RandomChoice: вариант %q никогда не выбирается", choice)
		}
	}
}

func BenchmarkRandomInt(b *testing.B) {
	InitRandom()
	for i := 0; i < b.N; i++ {
		RandomInt(1, 100)
	}
}

func BenchmarkRandomFloat(b *testing.B) {
	InitRandom()
	for i := 0; i < b.N; i++ {
		RandomFloat(0.0, 100.0)
	}
}

func BenchmarkRandomSlice(b *testing.B) {
	InitRandom()
	for i := 0; i < b.N; i++ {
		RandomSlice(10, 1, 100)
	}
}