package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInt(x, y string) (int, int) {
	var num1, num2 int
	num1, _ = strconv.Atoi(x)
	num2, _ = strconv.Atoi(y)
	return num1, num2
}

func RomanToInt(s string) int {
	var v, lv, cv int
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	for i := len(s) - 1; i >= 0; i-- {
		cv = h[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
	}

	return v
}

func IntToRoman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func Summ(a, b int) int {
	return a + b
}

func Min(a, b int) int {
	return a - b
}

func Umn(a, b int) int {
	return a * b
}

func Del(a, b int) int {
	return a / b
}

func OperazioneRoma(x, oper, y string) {
	inted := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	romans := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(inted); i++ {
		for j := 0; j < len(inted); j++ {
			if x == inted[i] && y == inted[j] {
				num1, num2 := ParseInt(x, y)
				switch oper {
				case "+":
					fmt.Println(Summ(num1, num2))
				case "-":
					fmt.Println(Min(num1, num2))
				case "*":
					fmt.Println(Umn(num1, num2))
				case "/":
					fmt.Println(Del(num1, num2))
				}
			}
			if x == romans[i] && y == romans[j] {
				num1 := RomanToInt(x)
				num2 := RomanToInt(y)
				switch oper {
				case "+":
					rim := Summ(num1, num2)
					fmt.Println(IntToRoman(rim))
				case "-":
					rim := Min(num1, num2)
					if rim < 1 {
						fmt.Print("Ошибка: в римской системе нет чисел меньше единицы.")
					}
					fmt.Println(IntToRoman(rim))
				case "*":
					rim := Umn(num1, num2)
					fmt.Println(IntToRoman(rim))
				case "/":
					rim := Del(num1, num2)
					fmt.Println(IntToRoman(rim))
				}
			}
			if (x == romans[i] && y == inted[j]) || (x == inted[i] && y == romans[j]) {
				fmt.Println("Ошибка: разные типы счисления.")
			}
		}
	}
}

func main() {
	var a, oper, b string
	var line string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line = sc.Text()
	viv := strings.Split(line, " ")
	if len(viv) != 3 {
		fmt.Println("Ошибка: так как формат матиматической операции не удовлетворяет заданию - два операнда и один оператор (+, -, /, *).")
		return
	}
	a, oper, b = viv[0], viv[1], viv[2]
	OperazioneRoma(a, oper, b)
}
