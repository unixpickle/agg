package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/unixpickle/essentials"
)

func main() {
	flag.Usage = func() {
		for _, line := range usageLines() {
			fmt.Fprintln(os.Stderr, line)
		}
		fmt.Fprintln(os.Stderr, "Optional flags:")
		flag.PrintDefaults()
	}

	var strictMode bool
	flag.BoolVar(&strictMode, "strict", false, "fail if an input token is invalid")
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	aggName := flag.Args()[0]

	agg, ok := Aggregates[aggName]
	if !ok {
		essentials.Die("unknown aggregate:", os.Args[1])
	}
	fmt.Println(agg(readFloats(strictMode)))
}

func readFloats(strictMode bool) <-chan float64 {
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
					if strictMode {
						essentials.Die("invalid token:", err)
					} else {
						fmt.Fprintln(os.Stderr, "skipping invalid token:", trimmedLine)
					}
				} else {
					res <- val
				}
			}
		}
	}()
	return res
}

func usageLines() []string {
	usage := []string{
		"Usage: " + os.Args[0] + " [flags] <aggregate type>",
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
