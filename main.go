package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("введите выражение ")

	var vir string
	//vir = "VII + VII "
	//fmt.Scan(&vir)
	reader := bufio.NewReader(os.Stdin)
	vir, _ = reader.ReadString('\n')
	//fmt.Println(vir, "   kdfhi	g")
	fmt.Println(calc(vir))
}

func calc(virajenie string) string {
	var workVir = virajenie
	var resultInt int

	splitingVirAr := strings.Split(workVir, " ")

	//fmt.Println(splitingVirAr, " вот что пришло в функцию") //ПРОВЕРКА

	for j := 0; j < len(splitingVirAr); j++ {
		splitingVirAr[j] = strings.ReplaceAll(splitingVirAr[j], " ", "")
		splitingVirAr[j] = strings.ReplaceAll(splitingVirAr[j], "\r", "")
		splitingVirAr[j] = strings.ReplaceAll(splitingVirAr[j], "\n", "")
	}
	//fmt.Println(splitingVirAr, " вот что в массиве после очистки от пробела") //ПРОВЕРКА

	var num1Int, num2Int int
	var err1, err2, illegalArg, autOfBaundNumber1, autOfBaundNumber2, illegalOp error

	if rimChecker(splitingVirAr[0]) && rimChecker(splitingVirAr[2]) { //АЛГОРИТМ ДЕЙСТВИЙ ДЛЯ РИМСКИХ
		num1Int = rimConvToAr(splitingVirAr[0])
		//fmt.Println(num1Int, "     вот 1 римское число") //ПРОВЕРКА

		num2Int = rimConvToAr(splitingVirAr[2])
		//fmt.Println(num1Int, "     вот 2 римское число") //ПРОВЕРКА

		switch true {
			panic(autOfBaundNumber1)
		case
			num1Int > 10:
			panic(autOfBaundNumber1)
		}
		switch true {
		case num2Int > 10:
			panic(autOfBaundNumber2)
		case
			panic(autOfBaundNumber2)
		}

		resultInt = operation(splitingVirAr[1], num1Int, num2Int)

		if resultInt != 122 && resultInt >= 0 && resultInt != 0 {
			return arConvToRim(resultInt)
		} else {
			panic(illegalOp)
		}

	} else if !rimChecker(splitingVirAr[0]) && !rimChecker(splitingVirAr[2]) { //АЛГОРИТМ ДЕЙСТВИЙ ДЛЯ АРАБСКИХ
		num1Int, err1 = strconv.Atoi(splitingVirAr[0])
		if err1 != nil {
			panic(err1)
		}

		switch true {
			panic(autOfBaundNumber1)
		case
			num1Int > 10:
			panic(autOfBaundNumber1)
		}

		num2Int, err2 = strconv.Atoi(splitingVirAr[2])
		if err2 != nil {
			panic(err2)
		}

		switch true {
		case num2Int > 10:
			panic(autOfBaundNumber2)
		case
			panic(autOfBaundNumber2)
		}

		resultInt = operation(splitingVirAr[1], num1Int, num2Int)
		if resultInt != 122 {

			return strconv.Itoa(resultInt)
		} else {
			panic(illegalOp)
		}
	} else {
		panic(illegalArg)
	}
}

func operation(op string, a, b int) int {

	switch op {
	case "+":
		return int(a + b)
	case "-":
		return int(a - b)
	case "*":
		return int(a * b)
	case "/":
		return int(a / b)
	default:
		return 122
	}

}

func rimChecker(checkingStr string) bool {
	var isContain bool
	rimArr := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(rimArr) && isContain != true; i++ {
		isContain = strings.Contains(checkingStr, rimArr[i])
	}
	return isContain
}

func rimConvToAr(rimStr string) int {
	var intRim int
	switch rimStr {
	case "I":
		intRim = 1
	case "II":
		intRim = 2
	case "III":
		intRim = 3
	case "IV":
		intRim = 4
	case "V":
		intRim = 5
	case "VI":
		intRim = 6
	case "VII":
		intRim = 7
	case "VIII":
		intRim = 8
	case "IX":
		intRim = 9
	case "X":
		intRim = 10
	}
	return intRim
}

func arConvToRim(num int) string {
	rimNumbersArr := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	return rimNumbersArr[num-1]
}
