package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ipbook := make(map[string]int)
	for scanner.Scan() {

		line := scanner.Text()
		n := strings.IndexByte(line, ' ')
		if n < 0 {
			continue
		}
		ip := line[:n]

		ipbook[ip]++
	}

	for k, v := range ipbook {
		fmt.Println(k, v)
	}
}
