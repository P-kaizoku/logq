package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type IPentry struct {
	ip    string
	count int
}

type StatusEntry struct {
	status string
	count  int
}

var ipbook = make(map[string]int)
var statusbook = make(map[string]int)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	field := flag.String("field", "ip", "ip | status")
	n := flag.Int("n", 10, "how many top entries")
	flag.Parse()

	for scanner.Scan() {
		line := scanner.Text()

		switch *field {
		case "ip":
			s, ok := ParseIPs(line)
			if !ok {
				continue
			}
			ipbook[s]++
		case "status":
			s, ok := ParseStatus(line)
			if !ok {
				continue
			}
			statusbook[s]++

		}
	}

	switch *field {
	case "ip":
		top(ipbook, *n)
	case "status":
		top(statusbook, *n)
	}

}

func ParseStatus(line string) (string, bool) {
	for f := range strings.FieldsSeq(line) {
		if len(f) == 3 && isDigits(f) {
			return f, true
		}
	}
	return "", false
}

func isDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func ParseIPs(line string) (string, bool) {
	n := strings.IndexByte(line, ' ')
	if n < 0 {
		return "", false
	}
	ip := line[:n]
	return ip, true

}

func top(m map[string]int, n int) {
	type entry struct {
		key   string
		count int
	}

	var top []entry

	for k, v := range m {
		top = append(top, entry{k, v})
	}
	sort.Slice(top, func(i, j int) bool { return top[i].count > top[j].count })
	for i := 0; i < n && i < len(top); i++ {
		fmt.Println(top[i].key, " --> ", top[i].count)
	}

}
