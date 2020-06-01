package main

import (
	"fmt"
	"log"
	"strconv"
)

func getNextDigit(item *string) int {
	var str string = *item
	var result, length int
	length = len(str)
	if len(*item) == 0 {
		return 0
	} else {
		num, err := strconv.Atoi(*item)
		if err != nil {
			log.Fatalln("ошибка приведения к числу:", err)
		}
		result = num % 10
	}
	*item = str[:length-1]
	return result
}

func getSum(a int, b int, mind int) (int, int) {
	if a+b+mind >= 10 {
		return (a + b + mind) % 10, 1
	} else {
		return a + b + mind, 0
	}
}

func main() {
	var item0, item1 string
	var num0, num1 int
	var result string
	fmt.Scan(&item0, &item1)
	var mind int
	for {
		num0 = getNextDigit(&item0)
		num1 = getNextDigit(&item1)
		sum, newMind := getSum(num0, num1, mind)
		mind = newMind
		result = strconv.Itoa(sum) + result
		if len(item0) == 0 && len(item1) == 0 && newMind == 0 {
			fmt.Println(result)
			return
		}

	}
}
