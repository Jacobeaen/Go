package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Есть ли пара [i, j] в массиве [[], [], ... , []]
func IsPairInArray(pair []int, array [][]int) bool {
	for _, brackets := range array {
		if brackets[0] == pair[0] && brackets[1] == pair[1] {
			return true
		}
	}

	return false
}

// Есть ли символ в строке
func IsSymbolInString(el rune, source []rune) bool {
	for _, x := range source {
		if x == el {
			return true
		}
	}

	return false
}

// Последний элемент массива
func GetLastElement(array []string) string {
	return array[len(array)-1]
}

// Проверка на правильность расставления скобок
func IsBracketsCorrect(source string) bool {
	count := 0
	for _, el := range source {
		if count < 0 {
			return false
		}

		if el == '(' {
			count++
		} else if el == ')' {
			count--
		}
	}

	return count == 0
}

// Является ли символ допустимым
func IsSymbolAllowed(symbol rune, array []rune) bool {
	for _, el := range array {
		if symbol == el {
			return true
		}
	}

	return false
}

// Проверка строки на правильный формат
func IsStringCorrect(source string) bool {
	array := []rune{
		'+', '0', '1',
		'-', '2', '3',
		'*', '4', '5',
		'/', '6', '7',
		'(', '8', '9',
		')', ' ', '.',
	}

	for _, el := range source {
		if !IsSymbolAllowed(el, array) {
			return false
		}
	}

	return IsBracketsCorrect(source)
}

// Слайс из индексов закрывающих скобок
func OpenBracketsIndexes(source string) (result []int, err error) {
	for index, el := range source {
		if el == '(' {
			result = append(result, index)
		}
	}

	if len(result) != 0 {
		return result, nil
	}

	return result, errors.New("there is no open brackets in string")
}

// Слайс из индексов закрывающих скобок
func CloseBracketsIndexes(source string) (result []int, err error) {
	for index, el := range source {
		if el == ')' {
			result = append(result, index)
		}
	}
	if len(result) != 0 {
		return result, nil
	}

	return result, errors.New("there is no close brackets in string")

}

// Кол-во элементов в срезе меньше указаного
func MinLength(element, start int, open []int) int {
	count := 0
	for _, open_br := range open[start:] {
		if element > open_br {
			count += 1
		} else {
			break
		}
	}

	return count
}

// Индексы открывающей и закрывающей ее скобок
func PairsBracketsIndexes(source string, open []int) ([][]int, error) {
	result := make([][]int, 0)

	if len(open) == 0 {
		return nil, errors.New("there is no brackets in the string")
	}

	for _, open_br := range open {
		close_br := open_br + 1
		count := 1

		for count != 0 {
			if source[close_br] == '(' {
				count++
			} else if source[close_br] == ')' {
				count--
			}
			close_br++
		}
		close_br--

		pair := []int{open_br, close_br}
		result = append(result, pair)
	}

	return result, nil
}

// Есть в строке скобки
func IsSubstringHaveBrackets(source string, pair []int) bool {
	start_i := pair[0] + 1
	end_i := pair[1]

	for _, symbol := range source[start_i:end_i] {
		if symbol == '(' || symbol == ')' {
			return false
		}
	}

	return true
}

// Список скобок без внутрених скобок
func OnlySimpleBreakets(source string, array [][]int) [][]int {
	result := make([][]int, 0)

	for _, pair := range array {
		if IsSubstringHaveBrackets(source, pair) {
			if !IsPairInArray(pair, result) {
				result = append(result, pair)
			}
		}
	}

	return result
}

// Операция над двумя числами
func Operation(array []float64, operation rune) (float64, error) {
	if len(array) < 2 {
		return 0, errors.New("wrong!")
	}

	number1 := array[len(array)-2]
	number2 := array[len(array)-1]

	if operation == '+' {
		return number1 + number2, nil
	} else if operation == '-' {
		return number1 - number2, nil
	} else if operation == '*' {
		return number1 * number2, nil
	} else if operation == '/' {
		if number2 != 0 {
			return number1 / number2, nil
		}
		return 0, errors.New("zero division error")
	}

	return 0, errors.New("unknow operation")
}

func IsSymbolDigit(symbol rune) bool {
	return symbol >= '0' && symbol <= '9'
}

func CalculateExpression(source string) (float64, error) {
	numbers := make([]float64, 0)
	operations := make([]rune, 0)
	signs := []rune{
		'+',
		'-',
		'*',
		'/',
	}
	rating := map[rune]int{
		'+': 1,
		'-': 1,
		'/': 2,
		'*': 2,
	}

	var last_digit string = "_"
	var last_sign rune = '_'
	var last_symbol rune = '_'

	count_numbers := 0
	count_operations := 0

	for _, symbol := range source {
		if IsSymbolDigit(symbol) || symbol == '.' {

			// Если еще не было символов
			if last_digit == "_" {
				numbers = append(numbers, float64(symbol-'0'))

				count_numbers++
				last_digit = string(symbol)

				// Если два или больше цифр подряд
			} else if IsSymbolDigit(last_symbol) || last_symbol == '.' {
				x := last_digit + string(symbol)

				n, _ := strconv.ParseFloat(x, 64)

				count_numbers--
				numbers[count_numbers] = float64(n)
				count_numbers++

				last_digit = x

				// Если просто одиночная цифра
			} else if symbol == '.' {
				last_digit = last_digit + string(symbol)

			} else {
				numbers = append(numbers, float64(symbol-'0'))

				last_digit = string(symbol)
				count_numbers++
			}

		} else if IsSymbolInString(symbol, signs) {
			if len(numbers) < 1 {
				return numbers[0], errors.New("wrong format")

				// Если первый знак - просто добавляем
			} else if len(operations) == 0 {
				operations = append(operations, symbol)

				count_operations++
				last_sign = symbol
			} else {
				if rating[last_sign] >= rating[symbol] {
					if len(numbers) < 2 {
						return 0, errors.New("wrong fromat!")
					}
					result, err := Operation(numbers, last_sign)
					if err != nil {
						return 0, errors.New("wrong fromat!")
					}

					count_numbers -= 2
					numbers = numbers[:count_numbers+1]
					numbers[count_numbers] = result
					count_numbers++

					count_operations--
					operations = operations[:count_operations]
					operations = append(operations, symbol)

					if len(numbers) > 1 && len(operations) > 1 {
						if rating[last_sign] >= rating[symbol] {
							result2, _ := Operation(numbers, operations[len(operations)-2])

							count_numbers -= 2
							numbers = numbers[:count_numbers+1]
							numbers[count_numbers] = result2
							count_numbers++

							operations = operations[count_operations:]
							count_operations--
						}
					}
					last_sign = symbol
					count_operations++

				} else {
					operations = append(operations, symbol)

					count_operations++
					last_sign = symbol
				}
			}
		}
		last_symbol = symbol
	}

	j := len(numbers) - 1
	for i := len(operations) - 1; i >= 0; i-- {
		result, err := Operation(numbers, operations[i])

		if err != nil {
			return 0, errors.New("wrong!")
		}
		numbers = numbers[:j]
		j--
		numbers[j] = result

		operations = operations[:i]
	}

	return numbers[0], nil

}

func Calc(source string) (float64, error) {
	if !IsStringCorrect(source) {
		return 0, errors.New("wrong format!")
	}

	source = strings.ReplaceAll(source, " ", "")
	for {
		x, _ := OpenBracketsIndexes(source)
		z, _ := PairsBracketsIndexes(source, x)
		arr := OnlySimpleBreakets(source, z)

		if len(arr) != 0 {
			for i := len(arr) - 1; i >= 0; i-- {
				start_br := arr[i][0]
				end_br := arr[i][1]

				substring := source[start_br+1 : end_br]
				result, _ := CalculateExpression(substring)

				substitution := strconv.FormatFloat(result, 'f', 2, 64)
				source = source[:start_br] + substitution + source[end_br+1:]
			}

		} else {
			result, err := CalculateExpression(source)

			if err != nil {
				return 0, errors.New("wrong!")
			}
			return result, nil
		}

		// fmt.Println(s)
		// fmt.Println(arr2)

	}
}

func main() {
	//s := "((1 + 2) * 3) / ((6 * 3) - ((2 * 3) + 9 + 2))"
	s := "5.5 - 3.3"
	//s := "(1 + 3 + 5) * 2.5 - 12"

	fmt.Print(Calc(s))

}
