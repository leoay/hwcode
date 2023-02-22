package main

import (
	"fmt"
	"strings"
)

func main() {
	testdata := []string{"10", "4,2,1", "9", "3,6,9,2", "6,3,4", "8"}
	n := len(testdata)
	//fmt.Scan(&n)
	var portgroup [][]string
	var portgrpmaparr []map[string]string

	for i := 0; i < n; i++ {
		//input := bufio.NewScanner(os.Stdin)
		//input.Scan()
		//instr := input.Text()
		instr := testdata[i]
		inarr := strings.Split(instr, ",")
		//将端口存储到map中
		portmap := make(map[string]string)
		for _, v := range inarr {
			if _, ok := portmap[v]; ok {
				//存在则忽略
			} else {
				portmap[v] = v
			}
		}
		if len(portgrpmaparr) > 0 {
			//比较当前map与历史map
			flag1 := true
			for k, v := range portgrpmaparr {
				flag := 0
				for _, v1 := range portmap {
					if _, ok := v[v1]; ok {
						flag++
					}
				}
				if flag >= 2 {
					flag = 0
					for _, v2 := range portmap {
						v[v2] = v2
					}
					portgrpmaparr[k] = v
				} else {
					flag1 = false
				}
			}
			if !flag1 {
				flag1 = true
				portgrpmaparr = append(portgrpmaparr, portmap)
			}
		} else {
			portgrpmaparr = append(portgrpmaparr, portmap)
		}
	}

	fmt.Println(portgrpmaparr)

	for _, v := range portgrpmaparr {
		var portarr []string
		for _, v1 := range v {
			portarr = append(portarr, v1)
		}
		portgroup = append(portgroup, portarr)
	}
	fmt.Println(portgroup)
}
