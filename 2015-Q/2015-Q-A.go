//+build ignore
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pivotal-golang/bytefmt"
	"github.com/zlowred/gcj"
	"runtime"
)

type test struct {
	num  int
	smax int
	s    []int
}

func solve(ch chan gcj.Result, t *test) {
	missing, total := 0, 0
	for i, x := range t.s {
		if total < i {
			missing += i - total
			total = i
		}
		total += x
	}
	time.Sleep(time.Second * 5)
	ch <- gcj.Result{t.num, fmt.Sprintf("%d", missing)}
	fmt.Printf("Case #%d is completed\n", t.num+1)
}

func runSolution() []string {
	T := gcj.NextInt()
	results := make([]string, T)
	defer gcj.Close(results)
	ch := make(chan gcj.Result, T)
	for i := 0; i < T; i++ {
		smax := gcj.NextInt()
		ss := []rune(gcj.NextWord())
		t := &test{i, smax, make([]int, smax+1)}
		for i := 0; i <= smax; i++ {
			t.s[i], _ = strconv.Atoi(string(ss[i]))
		}
		go solve(ch, t)
	}
	for i := 0; i < T; i++ {
		x := <-ch
		results[x.Num] = x.Res
	}
	return results
}

const testData = `4
4 11111
1 09
5 110011
0 1
`
const expectedResult = `Case #1: 0
Case #2: 1
Case #3: 2
Case #4: 0
`
const testName = "2015-Q-A"

const example = false

func verify(results []string) {
	if diff, err := gcj.VerifyTestData(expectedResult, results); err != nil {
		panic(err)
	} else if len(diff) > 0 {
		panic("\n" + diff)
	}
}

func main() {
	start := time.Now()
	defer func(t time.Time) {
		fmt.Printf("Processing took %v\n", time.Now().Sub(t))
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		fmt.Printf("Memory usage: %s\n", bytefmt.ByteSize(mem.Alloc))
	}(start)
	if example {
		gcj.SetTestData(testData)
	} else {
		gcj.SetName(testName)
	}
	results := runSolution()
	if example {
		verify(results)
	}
}
