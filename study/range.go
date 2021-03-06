//range（续）
//可以将下标或值赋予 _ 来忽略它。
//
//若你只需要索引，去掉 , value 的部分即可。

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i,_ := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
