package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/proc/net/ip_tables_names")
	if err != nil {
		defer file.Close()
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "NG")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "filter" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "OK")
			return
		} else {
			continue
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "NG")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
