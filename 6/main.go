package main

import (
	"example/random/utils"
	"fmt"
	"strings"
	"time"
)

// Цветовые коды для терминала
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	ColorBold   = "\033[1m"
)

func main() {
	// Инициализация генератора случайных чисел
	utils.InitRandom()

	// Красивый заголовок
	printHeader("🎲 Генератор Случайных Чисел 🎲")
	
	// Пример 1: Генерация случайных целых чисел
	printSection("1. Случайные целые числа", ColorBlue)
	fmt.Printf("   Бросок кубика (1-6): %s%d%s\n", 
		ColorYellow, utils.RandomInt(1, 6), ColorReset)
	fmt.Printf("   Случайный возраст (18-65): %s%d%s лет\n", 
		ColorYellow, utils.RandomInt(18, 65), ColorReset)
	fmt.Printf("   Случайный год (1990-2024): %s%d%s\n", 
		ColorYellow, utils.RandomInt(1990, 2024), ColorReset)
	
	// Пример 2: Генерация случайных дробных чисел
	printSection("2. Случайные дробные числа", ColorGreen)
	fmt.Printf("   Температура (18.0-28.0): %s%.2f%s°C\n", 
		ColorYellow, utils.RandomFloat(18.0, 28.0), ColorReset)
	fmt.Printf("   Процент скидки (5-25): %s%.1f%s%%\n", 
		ColorYellow, utils.RandomFloat(5.0, 25.0), ColorReset)
	fmt.Printf("   Коэффициент (0.1-1.0): %s%.3f%s\n", 
		ColorYellow, utils.RandomFloat(0.1, 1.0), ColorReset)
	
	// Пример 3: Генерация массива случайных чисел
	printSection("3. Массив случайных чисел", ColorPurple)
	lotteryNumbers := utils.RandomSlice(6, 1, 49)
	fmt.Print("   Лотерейные числа: ")
	for i, num := range lotteryNumbers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s%02d%s", ColorCyan, num, ColorReset)
	}
	fmt.Println()
	
	scores := utils.RandomSlice(5, 60, 100)
	fmt.Print("   Оценки студентов: ")
	for i, score := range scores {
		if i > 0 {
			fmt.Print(" | ")
		}
		color := getScoreColor(score)
		fmt.Printf("%s%d%s", color, score, ColorReset)
	}
	fmt.Println()
	
	// Пример 4: Генерация булевых значений
	printSection("4. Случайные булевы значения", ColorCyan)
	for i := 0; i < 3; i++ {
		result := utils.RandomBool()
		if result {
			fmt.Printf("   Попытка %d: %s✅ Успех!%s\n", i+1, ColorGreen, ColorReset)
		} else {
			fmt.Printf("   Попытка %d: %s❌ Неудача%s\n", i+1, ColorRed, ColorReset)
		}
	}
	
	// Пример 5: Случайный выбор из списка
	printSection("5. Случайный выбор из списка", ColorYellow)
	
	// Выбор цвета
	colors := []string{"🔴 Красный", "🟢 Зелёный", "🔵 Синий", "🟡 Жёлтый", "🟣 Фиолетовый"}
	fmt.Printf("   Выбранный цвет: %s%s%s\n", 
		ColorBold, utils.RandomChoice(colors), ColorReset)
	
	// Выбор действия
	actions := []string{"⚔️ Атака", "🛡️ Защита", "✨ Магия", "🏃 Побег"}
	fmt.Printf("   Действие героя: %s%s%s\n", 
		ColorBold, utils.RandomChoice(actions), ColorReset)
	
	// Выбор направления
	directions := []string{"⬆️ Север", "➡️ Восток", "⬇️ Юг", "⬅️ Запад"}
	fmt.Printf("   Направление движения: %s%s%s\n", 
		ColorBold, utils.RandomChoice(directions), ColorReset)
	
	// Пример 6: Игровой пример - бросок монеты
	printSection("6. Игра: Орёл или Решка", ColorRed)
	fmt.Println("   Подбрасываем монету...")
	time.Sleep(500 * time.Millisecond)
	
	for i := 0; i < 5; i++ {
		if utils.RandomBool() {
			fmt.Printf("   Бросок %d: %s🦅 ОРЁЛ%s\n", i+1, ColorYellow, ColorReset)
		} else {
			fmt.Printf("   Бросок %d: %s👑 РЕШКА%s\n", i+1, ColorCyan, ColorReset)
		}
		time.Sleep(300 * time.Millisecond)
	}
	
	// Пример 7: Генерация пароля
	printSection("7. Генератор паролей", ColorGreen)
	password := generatePassword(12)
	fmt.Printf("   Сгенерированный пароль: %s%s%s\n", 
		ColorBold+ColorYellow, password, ColorReset)
	
	// Пример 8: Статистика
	printSection("8. Статистика случайных чисел", ColorBlue)
	showStatistics()
	
	// Завершение
	printFooter()
}

// Вспомогательные функции для красивого вывода

func printHeader(title string) {
	width := 50
	fmt.Println()
	fmt.Printf("%s%s%s\n", ColorBold+ColorCyan, strings.Repeat("═", width), ColorReset)
	// Используем простой центрированный вывод без подсчёта длины
	fmt.Printf("%s%*s%s\n", ColorBold+ColorCyan, 
		(width+len(title))/2, title, ColorReset)
	fmt.Printf("%s%s%s\n", ColorBold+ColorCyan, strings.Repeat("═", width), ColorReset)
	fmt.Println()
}

func printSection(title string, color string) {
	fmt.Printf("\n%s%s%s %s%s\n", ColorBold, color, "▶", title, ColorReset)
	fmt.Printf("%s%s%s\n", color, strings.Repeat("─", 40), ColorReset)
}

func printFooter() {
	width := 50
	fmt.Println()
	fmt.Printf("%s%s%s\n", ColorCyan, strings.Repeat("═", width), ColorReset)
	fmt.Printf("%s%s%s%s\n", ColorCyan, 
		strings.Repeat(" ", 15), "✨ Готово! ✨", ColorReset)
	fmt.Printf("%s%s%s\n", ColorCyan, strings.Repeat("═", width), ColorReset)
	fmt.Println()
}

func getScoreColor(score int) string {
	switch {
	case score >= 90:
		return ColorGreen
	case score >= 75:
		return ColorYellow
	case score >= 60:
		return ColorBlue
	default:
		return ColorRed
	}
}

func generatePassword(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	password := ""
	for i := 0; i < length; i++ {
		idx := utils.RandomInt(0, len(chars)-1)
		password += string(chars[idx])
	}
	return password
}

func showStatistics() {
	// Генерируем 100 случайных чисел и считаем статистику
	numbers := utils.RandomSlice(100, 1, 10)
	
	// Подсчёт частоты
	frequency := make(map[int]int)
	sum := 0
	for _, num := range numbers {
		frequency[num]++
		sum += num
	}
	
	// Вывод гистограммы
	fmt.Println("   Распределение чисел от 1 до 10 (100 бросков):")
	for i := 1; i <= 10; i++ {
		count := frequency[i]
		bar := strings.Repeat("█", count/2)
		if count%2 == 1 {
			bar += "▌"
		}
		fmt.Printf("   %2d: %s%s%s %d\n", i, ColorCyan, bar, ColorReset, count)
	}
	
	avg := float64(sum) / float64(len(numbers))
	fmt.Printf("   Среднее значение: %s%.2f%s\n", ColorYellow, avg, ColorReset)
}