package dup

import (
	"bufio"
	"fmt"
	"os"
)
func Dup1(){
	counts := make(map[string] int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++

	}

	// 忽视Input.Err()中可能的错误
	for line, n :=range counts{
		if n>1 {
			fmt.Println("%d\t%s\n", n, line)

		}
	}
}