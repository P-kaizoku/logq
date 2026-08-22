package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	paths := []string{"/api/users", "/api/login", "/health", "/api/orders", "/static/app.css"}
	statuses := []int{200, 200, 200, 200, 200, 301, 401, 404, 500}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for range 4_000_000 {
		ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		fmt.Fprintf(w, "%s - - [22/Aug/2026:10:%02d:%02d] \"GET %s HTTP/1.1\" %d %d \"-\" \"Mozilla/5.0\"\n",
			ip, rand.Intn(60), rand.Intn(60),
			paths[rand.Intn(len(paths))],
			statuses[rand.Intn(len(statuses))],
			rand.Intn(5000))
	}
}
