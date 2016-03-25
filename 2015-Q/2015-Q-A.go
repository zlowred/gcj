package main

import (
	"fmt"
	"github.com/zlowred/gcj"
	"strconv"
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
	ch <- gcj.Result{t.num, fmt.Sprintf("%d", missing)}
}

func main() {
	gcj.SetName("2015-Q-A-test")
	T := gcj.NextInt()
	results := make([]string, T)
	defer gcj.Close(results)
	ch := make(chan gcj.Result, T)
	defer close(ch)
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
}
