package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//计算至少需要多少个主快递站点

func main() {
	n := 0
	fmt.Scan(&n)
	var ss [][]string
	s2 := make()
	for i := 0; i < n; i++ {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		strrst := strings.Fields(input.Text())
		var s1 []string
		for _, v := range strrst {
			s1 = append(s1, v)
		}
		ss = append(ss, s1)
	}

	fmt.Println(ss)
}
