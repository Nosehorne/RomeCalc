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
		result, err := calculate(input)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			return
		}
		fmt.Println(result)
	}
}

func calculate(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", fmt.Errorf("неверный формат ввода")
	}

	a, aIsRoman, err := parseNumber(parts[0])
	if err != nil {
		return "", err
	}

	b, bIsRoman, err := parseNumber(parts[2])
	if err != nil {
		return "", err
	}

	if aIsRoman != bIsRoman {
		return "", fmt.Errorf("используются разные системы счисления")
	}

	op := parts[1]
	result, err := performOperation(a, b, op)
	if err != nil {
		return "", err
	}

	if aIsRoman {
		if result < 1 {
			return "", fmt.Errorf("результат работы с римскими числами должен быть больше нуля")
		}
		return intToRoman(result), nil
	}

	return fmt.Sprintf("%d", result), nil
}

func parseNumber(s string) (int, bool, error) {
	if val, ok := romanNumerals[s]; ok {
		return val, true, nil
	}

	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil || n < 1 || n > 10 {
		return 0, false, fmt.Errorf("недопустимое число: %s", s)
	}
	return n, false, nil
}

func performOperation(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неподдерживаемая операция: %s", op)
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
