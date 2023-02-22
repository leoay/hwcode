package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//【华为OD机试真题2023 Golang】 打印文件

type printTask struct {
	Num int
	Pro int
}

func insertSort(printtask []printTask, pro printTask) []printTask {
	k := len(printtask)
	a := 0
	for i := 0; i < k; i++ {
		if pro.Pro > printtask[i].Pro {
			a = i
			break
		}
	}
	printtask = append(printtask, pro)
	for j := k; j > a; j-- {
		printtask[j] = printtask[j-1]
	}
	printtask[a] = pro
	return printtask
}

func main() {
	n := 0
	var rst []string
	fmt.Scan(&n)
	pq := make(map[string][]printTask)
	num := 0
	for i := 0; i < n; i++ {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		instr := input.Text()
		inarry := strings.Fields(instr)
		pn := inarry[1]
		po := 0
		if len(inarry) > 2 {
			po, _ = strconv.Atoi(inarry[2])
		}

		if inarry[0] == "IN" {
			num++
			if val, ok := pq[pn]; ok {
				printmaking := printTask{
					Num: num,
					Pro: po,
				}
				val = insertSort(val, printmaking)
				pq[pn] = val
			} else {
				printmaking := printTask{
					Num: num,
					Pro: po,
				}
				var printList []printTask
				printList = append(printList, printmaking)
				pq[pn] = printList
			}
		}
		if inarry[0] == "OUT" {
			if val, ok := pq[pn]; ok {
				rst = append(rst, strconv.Itoa(val[0].Num))
				if len(pq[pn]) > 1 {
					pq[pn] = val[1:]
				} else {
					delete(pq, pn)
				}
			} else {
				rst = append(rst, "NULL")
			}
		}
	}
	for _, v := range rst {
		fmt.Println(v)
	}
}
