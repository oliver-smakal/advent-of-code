package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var text string
	var err error
	for err == nil {
		text, err = in.ReadString('\n')
		fmt.Print("foo:", text)
	}
}
