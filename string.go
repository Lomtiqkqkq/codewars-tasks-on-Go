package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {

}
func InsertTab() {
	var a string
	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	} else {
		fmt.Print(strings.Join(strings.Split(a, ""), "*"))
	}
}

// TODO разобраться в 33
func SqrtDigitToNumber() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(s); i++ {
		tmp := s[i] - '0' // или '48'
		fmt.Print(tmp * tmp)
	}
}
func exceptionError() (string, int) {
	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	result, err1 := divide(a, b)
	if err != nil || err1 != nil {
		fmt.Println("Ошибка")
		return "Ошибка", 0
	} else {
		fmt.Println(result)
		return "", result
	}
}
func divide(a int, b int) (int, error) {
	return a / b, nil
}
func validateRegExp() string {
	var a string
	fmt.Scan(&a)
	runes := []rune(a)
	reg := regexp.MustCompile(`^[a-zA-Z0-9]{5,}$`)
	if len(runes) >= 5 && reg.MatchString(string(runes)) {
		fmt.Print("OK")
		return "Ok"
	} else {
		fmt.Print("Wrong password")
		return "Wrong password"
	}
}
func deleteDuplicatesAll() string {
	var s string
	var result []string
	fmt.Scan(&s)
	runes := []rune(s)
	for i, _ := range runes {
		if strings.Count(string(runes), string(runes[i])) == 1 {
			result = append(result, string(runes[i]))
		}
	}
	fmt.Print(strings.Join(result, ""))
	return strings.Join(result, "")
}
func filterByOdd() {
	var s string
	fmt.Scan(&s)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if i%2 != 0 {
			fmt.Printf("%s", string(runes[i]))
		}
	}
}
func isRight() string {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace("\n")
	runes := []rune(text)
	if unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.' {
		fmt.Print("Right")
		return "Right"
	} else {
		fmt.Print("Wrong")
		return "Wrong"
	}
}
func isPalindrom() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeText := []rune(text)
	lenRune := len(runeText) - 2
	middle := lenRune / 2
	for i := 0; i != middle; i++ {
		if string(runeText[i]) == string(runeText[lenRune-i-1]) {
			continue
		} else {
			fmt.Println("Нет")
			return
		}
	}
	fmt.Print("Палиндром")
}
func idx() {
	var s, idx string
	fmt.Scanf("%s \n %s", &s, &idx)
	fmt.Print(strings.Index(s, idx))

}
func testWithRune() {
	var txt string = "Слово"
	txtRunes := []rune(txt)
	fmt.Printf("\nСрез рун: %v", txtRunes)
	fmt.Printf("\nИз рун получаем снова текст: %v", string(txtRunes))
	lenTxtRunes := len(txtRunes)
	fmt.Printf("\nДлина текста (среза): %v", lenTxtRunes)

	for i, v := range txtRunes {
		fmt.Printf("\nСимвол %d: %s", i, string(v))
	}

	txtRunes[lenTxtRunes-1] = rune('а')
	fmt.Printf("\nЗаменили последнюю букву на \"а\": %v", string(txtRunes))

	txtRunes = append(txtRunes[:4])
	fmt.Printf("\nОтрезали последнюю букву: %v", string(txtRunes))

	txtRunes = append(txtRunes, 'а', 'р', 'ь')
	fmt.Printf("\nДобавили три буквы в конец: %v", string(txtRunes))
}
