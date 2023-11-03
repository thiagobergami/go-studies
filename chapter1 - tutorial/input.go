package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/* var s, sep string
	fmt.Println(os.Args[0]) */
	/* for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		s += sep + os.Args[i]
		sep = " "
	} */

	/* for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	} */
	/* fmt.Println(len(s))
	fmt.Println(s[1:])
	fmt.Println(s) */

	fmt.Println(strings.Join(os.Args[1:], " "))
}
