//我们可以看出切片，实际的是获取数组的某一部分，len切片<=cap切片<=len数组，切片由三部分组成：指向底层数组的指针、len、cap。

package main
import "fmt"

func main() {
s := []int{2, 3, 5, 7, 11, 13}
printSlice(s)

// Slice the slice to give it zero length.
s = s[:0]
printSlice(s)

// Extend its length.
s = s[:5]
printSlice(s)

// Drop its first two values.
s = s[4:]
printSlice(s)
}

func printSlice(s []int) {
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}