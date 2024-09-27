package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	// Удаляем пробелы
	input = strings.ReplaceAll(input, " ", "")

	// Проверяем, есть ли в строке оператор
	operator, index := findOperator(input)
	if operator == "" {
		panic("Некорректный формат: строка не является математической операцией")
	}

	operand1 := input[:index]
	operand2 := input[index+1:]

	if operand1 == "" || operand2 == "" {
		panic("Некорректный формат: отсутствует операнд")
	}

	// Определяем систему счисления
	isRoman := isRomanNumeral(operand1) && isRomanNumeral(operand2)
	isArabic := isArabicNumeral(operand1) && isArabicNumeral(operand2)

	if !(isRoman || isArabic) {
		panic("Операнды должны быть оба римскими или оба арабскими числами")
	}

	if isRoman {
		a := romanToInt(operand1)
		b := romanToInt(operand2)
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("Римские числа должны быть от I до X")
		}
		result := calculate(a, b, operator)
		if result < 1 {
			panic("В римской системе нет чисел меньше I")
		}
		fmt.Println(intToRoman(result))
	} else {
		a, _ := strconv.Atoi(operand1)
		b, _ := strconv.Atoi(operand2)
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("Числа должны быть от 1 до 10")
		}
		result := calculate(a, b, operator)
		fmt.Println(result)
	}
}

// Функция для поиска оператора
func findOperator(input string) (string, int) {
	for i, char := range input {
		if strings.ContainsRune("+-*/", char) {
			return string(char), i
		}
	}
	return "", -1
}

// Функция для проверки арабского числа
func isArabicNumeral(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Функция для проверки римского числа
func isRomanNumeral(s string) bool {
	_, exists := romanMap[s]
	return exists
}

// Карта римских цифр
var romanMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

// Функция для преобразования римского числа в целое
func romanToInt(s string) int {
	return romanMap[s]
}

// Функция для преобразования целого числа в римское
func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result strings.Builder
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			result.WriteString(syms[i])
		}
	}
	return result.String()
}

// Функция для выполнения арифметической операции
func calculate(a int, b int, operator string) int {
	switch operator {
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
		panic("Некорректный оператор")
	}
}
