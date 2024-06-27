package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите выражение: ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "" {
			continue
		}
		result := calculate(input)
		fmt.Println(result)
	}
}

func calculate(input string) string {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}

	a, aIsRoman := parseNumber(parts[0])
	b, bIsRoman := parseNumber(parts[2])

	if aIsRoman != bIsRoman {
		panic("Используются разные системы счисления")
	}

	result := performOperation(a, b, parts[1])

	if aIsRoman {
		if result < 1 {
			panic("Результат работы с римскими числами должен быть больше нуля")
		}
		return intToRoman(result)
	}

	return fmt.Sprintf("%d", result)
}

func parseNumber(s string) (int, bool) {
	if val, ok := romanNumerals[s]; ok {
		return val, true
	}

	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil || n < 1 || n > 10 {
		panic("Недопустимое число: " + s)
	}
	return n, false
}

func performOperation(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль")
		}
		return a / b
	default:
		panic("Неподдерживаемая операция: " + op)
	}
}

func intToRoman(num int) string {
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}
	return roman
}
