package gcj

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
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

func Printf(format string, a ...interface{}) {
	if writer != nil {
		fmt.Fprintf(writer, format, a...)
	} else {
		fmt.Printf(format, a...)
	}
}
func Close(results []string) {
	for i, res := range results {
		Printf("Case #%d: %s\n", i+1, res)
	}

	if writer != nil {
		writer.Flush()
	}
}

func SetTestData(data string) {
	lineScanner = bufio.NewScanner(strings.NewReader(data))
	wordScanner = bufio.NewScanner(strings.NewReader(""))
	wordScanner.Split(bufio.ScanWords)
}
func VerifyTestData(expectedResult string, actualResult []string) (string, error) {
	actual := ""
	for i, res := range actualResult {
		actual += fmt.Sprintf("Case #%d: %s\n", i+1, res)
	}
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(expectedResult),
		B:        difflib.SplitLines(actual),
		FromFile: "Expected",
		ToFile:   "Actual",
		Context:  0,
		Eol:      "\n",
	}
	return difflib.GetUnifiedDiffString(diff)
}

func SetName(name string) {
	outputFile, err := os.Create(name + ".out")
	if err != nil {
		panic(err)
	}
	writer = bufio.NewWriter(outputFile)
	inputFile, err := os.Open(name + ".in")
	if err != nil {
		panic(err)
	}
	lineScanner = bufio.NewScanner(inputFile)
	wordScanner = bufio.NewScanner(strings.NewReader(""))
	wordScanner.Split(bufio.ScanWords)
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
