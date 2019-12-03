package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileNames := make(map[string]map[string]bool)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex-1.4: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			keys := make([]string, 0, len(fileNames[line]))
			for k := range fileNames[line] {
				keys = append(keys, k)
			}
			fmt.Println(strings.Join(keys, " "))
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if fileNames[input.Text()] == nil {
			fileNames[input.Text()] = make(map[string]bool)
		}
		fileNames[input.Text()][f.Name()] = true
	}
}
