package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func findUnique(input io.Reader, output io.Writer) {
	s := bufio.NewScanner(input)
	var result, zeroCount int
	for s.Scan() {
		val, _ := strconv.Atoi(s.Text())
		if val == 0 {
			zeroCount++
		}
		result = result ^ val
	}
	w := bufio.NewWriter(output)
	res := strconv.Itoa(result)
	if zeroCount > 0 && zeroCount%2 != 0 {
		w.WriteString("0")
	} else {
		w.WriteString(res)
	}
	w.Flush()
}

func main() {
	findUnique(os.Stdin, os.Stdout)
}
