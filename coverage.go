package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	covFile := "/devtools-coverage.out"
	readmeFile := os.Getenv("README_FILE")

	minCov, err := strconv.ParseInt(os.Getenv("MINIMUM_COVERAGE"), 10, 32)
	if err != nil {
		panic(err)
	}

	// read cov file line by line

	cov, err := os.Open(covFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = cov.Close()
	}()

	fmt.Println("coverage file opened")

	r := regexp.MustCompile("^total.*")

	var total string

	scanner := bufio.NewScanner(cov)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if r.Match([]byte(line)) {
			rr := regexp.MustCompile("\\s+")
			newline := rr.ReplaceAllString(line, " ")

			chunks := strings.Split(newline, " ")
			total = chunks[len(chunks)-1]
		}
	}

	if total == "" {
		panic("cannot find total")
	}

	fmt.Printf("coverage total is %s\n", total)

	tt := strings.ReplaceAll(total, "%", "")

	chunks := strings.Split(tt, ".")
	tt = chunks[0]

	t, err := strconv.ParseInt(tt, 10, 32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("coverage total integer is %d\n", t)

	var color string
	if t < 50 {
		color = "red"
	} else if t < 70 {
		color = "yellow"
	} else if t < 80 {
		color = "yellowgreen"
	} else if t < 90 {
		color = "green"
	} else {
		color = "brightgreen"
	}

	fmt.Printf("coverage color is %s\n", color)

	// update readme file

	cc, err := ioutil.ReadFile(readmeFile)
	if err != nil {
		panic(err)
	}

	newBadge := fmt.Sprintf(
		"![coverage-badge-do-not-edit](https://img.shields.io/badge/Coverage-%d%%25-%s.svg?longCache=true&style=flat)",
		t,
		color,
	)

	rrr := regexp.MustCompile(`!\[coverage-badge-do-not-edit.*`)

	newReadme := rrr.ReplaceAll(cc, []byte(newBadge))

	if err := ioutil.WriteFile(readmeFile, newReadme, 0); err != nil {
		panic(err)
	}

	fmt.Println("readme file updated with the new badge")

	if t < minCov {
		fmt.Printf("coverage is %d < %d\n", t, minCov)
		os.Exit(1)
	} else {
		fmt.Printf("OK coverage is %d >= %d\n", t, minCov)
	}
}
