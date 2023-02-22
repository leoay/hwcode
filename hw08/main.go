package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//【华为OD机试真题2023 Golang】 打印文件

func isSameChar(str1, str2 string) bool {
	for _, v := range str2 {
		flag := strings.Contains(str1, string(v))
		if flag {
			return false
		}
	}
	return true
}

func main() {
	//n := 0
	//var rst []string
	//fmt.Scan(&n)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	inputstr1 := input.Text()

	input = bufio.NewScanner(os.Stdin)
	input.Scan()
	inputstr2 := input.Text()

	fmt.Println(inputstr1, inputstr2)

	inputarr1 := strings.Fields(inputstr1)
	inputarr2 := strings.Fields(inputstr2)

	fmt.Println(inputarr1)

	fmt.Println(inputarr2)

	for _, v1 := range inputarr1 {
		for _, v2 := range inputarr2 {
			if len(v1) == len(v2) {
				flag := isSameChar(v1, v2)
				if flag {
					fmt.Println(true)
					return
				}
				continue
			}
		}
	}
	fmt.Println(false)
}
