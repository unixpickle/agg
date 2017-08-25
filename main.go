package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/unixpickle/essentials"
)

func main() {
	if len(os.Args) != 2 {
		for _, line := range usageLines() {
			fmt.Fprintln(os.Stderr, line)
		}
		os.Exit(1)
	}

	agg, ok := Aggregates[os.Args[1]]
	if !ok {
		essentials.Die("unknown aggregate:", os.Args[1])
	}
	fmt.Println(agg(readFloats()))
}

func readFloats() <-chan float64 {
	res := make(chan float64, 128)
	go func() {
		defer close(res)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			trimmedLine := strings.TrimSpace(scanner.Text())
			if len(trimmedLine) > 0 {
				val, err := strconv.ParseFloat(trimmedLine, 64)
				if err != nil {
					essentials.Die("bad line:", err)
				}
				res <- val
			}
		}
	}()
	return res
}

func usageLines() []string {
	usage := []string{
		"Usage: " + os.Args[0] + " <aggregate type>",
		"",
		"Available aggregate types:",
	}

	var aggNames []string
	var longest int
	for name := range AggregateUsage {
		aggNames = append(aggNames, name)
		if len(name) > longest {
			longest = len(name)
		}
	}
	sort.Strings(aggNames)
	for _, name := range aggNames {
		desc := AggregateUsage[name]
		for len(name) < longest {
			name += " "
		}
		usage = append(usage, "  "+name+"    "+desc)
	}

	// Ending with a blank line looks cleaner.
	usage = append(usage, "")

	return usage
}
