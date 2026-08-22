package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		// line := scanner.Text()
		// n := strings.IndexByte(line, ' ')
		// ip := line[:n]

		if len(fields) == 0 {
			continue
		}
		fmt.Println(fields[0])

		// fmt.Println(ip)
	}
}
