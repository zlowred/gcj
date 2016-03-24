package gcj

import (
	"bufio"
	"strings"
	"os"
	"strconv"
	"errors"
	"fmt"
)

type Result struct {
	Num int
	Res string
}

func init() {
	lineScanner = bufio.NewScanner(os.Stdin)
	wordScanner = bufio.NewScanner(strings.NewReader(""))
	wordScanner.Split(bufio.ScanWords)
}

var lineScanner *bufio.Scanner
var wordScanner *bufio.Scanner

var writer *bufio.Writer

func Printf(format string, a ... interface{}) {
	fmt.Printf(format, a ...)
	if writer != nil {
		fmt.Fprintf(writer, format, a ...)
	}
}
func Close() {
	if writer != nil {
		writer.Flush()
	}
}

func SetName(name string) {
	outputFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	writer = bufio.NewWriter(outputFile)
}

func NextLine() (string, error) {
	if !lineScanner.Scan() {
		if err := lineScanner.Err(); err != nil {
			panic(err)
		}
		return "", errors.New("nextLine: EOF reached")
	}
	return lineScanner.Text(), nil
}

func NextWord() string {
	for !wordScanner.Scan() {
		if err := wordScanner.Err(); err != nil {
			panic(err)
		}
		line, err := NextLine()
		if err != nil {
			return ""
		}
		wordScanner = bufio.NewScanner(strings.NewReader(line))
		wordScanner.Split(bufio.ScanWords)
	}
	return wordScanner.Text()
}

func NextInt() int {
	res, err := strconv.Atoi(NextWord())
	if err != nil {
		panic(err)
	}
	return res
}

func NextInt64() int64 {
	res, err := strconv.ParseInt(NextWord(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func NextFloat64() float64 {
	res, err := strconv.ParseFloat(NextWord(), 64)
	if err != nil {
		panic(err)
	}
	return res
}
