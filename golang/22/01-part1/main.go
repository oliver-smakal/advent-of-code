package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err         error
		maxCalories int
		calories    int
	)

	in := bufio.NewReader(os.Stdin)

	for err == nil {
		calories, err = readBlock(in)
		if calories > maxCalories {
			maxCalories = calories
		}
	}

	fmt.Println(maxCalories)
}

func readBlock(in *bufio.Reader) (int, error) {
	var (
		err  error
		sum  int
		line string
	)

	for err == nil {
		line, err = in.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil || line == "" {
			break
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			return sum, err
		}

		sum += val
	}

	return sum, err
}
