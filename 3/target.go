package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func printResult(w *bufio.Writer, hasCouple bool) {
	var result string
	if hasCouple == true {
		result = "1"
	} else {
		result = "0"
	}
	_, err := w.WriteString(result)
	if err != nil {
		log.Fatalln("ошибка записи файла: ", err)
	}
	w.Flush()
}

func parseInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalln("ошибка преобразования", err)
	}
	return result
}

func main() {
	inputFile, errInputFile := os.Open("input.txt")
	if errInputFile != nil {
		log.Fatalln("ошибка открытия файла: ", errInputFile)
	}
	defer inputFile.Close()
	outputFile, errOutputFile := os.OpenFile("output.txt", os.O_WRONLY, 0)
	if errOutputFile != nil {
		outputFile, _ = os.Create("output.txt")
	}
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	sumStr, errSum := reader.ReadString('\n')
	if errSum != nil {
		log.Fatalln("ошибка чтения суммы", errSum)
	}
	var sum int = parseInt(strings.Trim(sumStr, "\n"))
	h := make(map[int]int)
	for {
		value, errValue := reader.ReadString(' ')
		if errValue != nil && errValue != io.EOF {
			log.Fatalln("ошибка чтения последовательности", errValue)
		}
		var v string = strings.Trim(value, " ")
		if v == "" {
			if errValue == io.EOF {
				printResult(writer, false)
				return
			}
			continue
		}
		var item int = parseInt(v)
		if item >= sum {
			continue
		}
		if _, inHash := h[item]; inHash {
			printResult(writer, true)
			return
		}
		h[sum-item] = item
		if errValue == io.EOF {
			printResult(writer, false)
			return
		}
	}
}
