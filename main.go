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
		TopNIps(*n)
	case "status":
		TopStatus(*n)
	}

}

func ParseStatus(line string) (string, bool) {
	n := strings.LastIndexByte(line, '-')
	if n < 0 {
		return "", false
	}
	line = line[:n-1]

	i2 := strings.LastIndexByte(line, ' ')
	if i2 < 0 {
		return "", false
	}
	i1 := strings.LastIndexByte(line[:i2], ' ')
	if i1 < 0 {
		return "", false
	}

	status := line[i1+1 : i2]

	return status, true

}

func ParseIPs(line string) (string, bool) {
	n := strings.IndexByte(line, ' ')
	if n < 0 {
		return "", false
	}
	ip := line[:n]
	return ip, true

}

func TopNIps(n int) {
	var top []IPentry
	for k, v := range ipbook {
		top = append(top, IPentry{ip: k, count: v})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].count > top[j].count
	})

	for i := 0; i < n && i < len(top); i++ {
		fmt.Println(top[i].ip, " --> ", top[i].count)
	}
}

func TopStatus(n int) {
	var top []StatusEntry
	for k, v := range statusbook {
		top = append(top, StatusEntry{status: k, count: v})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].count > top[j].count
	})

	for i := 0; i < n && i < len(top); i++ {
		fmt.Println(top[i].status, " --> ", top[i].count)
	}

}
