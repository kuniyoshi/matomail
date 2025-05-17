package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	helpFlag := flag.Bool("help", false, "Display usage information")
	regexPattern := flag.String("pattern", "", "Regular expression pattern for line comparison masking")
	flag.Parse()

	if *helpFlag {
		displayHelp()
		return
	}

	var re *regexp.Regexp
	if *regexPattern != "" {
		var err error
		re, err = regexp.Compile(*regexPattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid regular expression pattern: %v\n", err)
			os.Exit(1)
		}
	}

	stdout := bufio.NewWriterSize(os.Stdout, 1)
	scanner := bufio.NewScanner(os.Stdin)

	var prevLine string
	var currentLine string
	var count int

	for scanner.Scan() {
		prevLine = currentLine
		currentLine = scanner.Text()

		isSame := areLinesSame(currentLine, prevLine, re)

		fmt.Fprintf(os.Stderr, "DEBUG: isSame: %v\n", isSame)

		if !isSame {
			count = 0
			fmt.Fprintf(stdout, "%s\n", currentLine)
			stdout.Flush()
			continue
		}

		count++

		if count == 1 {
			fmt.Fprint(stdout, "\033[1A")
			fmt.Fprint(stdout, "\033[2K")
			fmt.Fprintf(stdout, "(1) %s\n", prevLine)
			fmt.Fprintf(stdout, "(2) %s\n", currentLine)
			stdout.Flush()
		} else {
			fmt.Fprint(stdout, "\033[1A")
			fmt.Fprint(stdout, "\033[2K")
			fmt.Fprintf(stdout, "...\n")
			fmt.Fprintf(stdout, "(%d) %s\n", count+1, currentLine)
			stdout.Flush()
		}
	}
}

func areLinesSame(currentLine, prevLine string, re *regexp.Regexp) bool {
	if re == nil {
		return currentLine == prevLine
	}
	return re.ReplaceAllString(currentLine, "") == re.ReplaceAllString(prevLine, "")
}

	// エラーチェック
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}
}

// 使用方法を表示する関数
func displayHelp() {
	helpText := `
matomail - Combines "matome" (gather/collect) with "tail"

USAGE:
  tail /var/log/foo.log | matomail [OPTIONS]

OPTIONS:
  --help               Display this help message
  --pattern=REGEX      Specify a regular expression pattern for line comparison masking
                       Example: --pattern="\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}"

DESCRIPTION:
  matomail reads lines from standard input and detects consecutive identical lines.
  When consecutive identical lines are found, it displays a counter prefix (e.g., "(3)")
  followed by the line content.

  By default, lines are considered identical when they match exactly. If a regular
  expression pattern is specified with --pattern, matching portions are masked during
  comparison, allowing lines with varying timestamps or other dynamic content to be
  treated as identical.
`
	fmt.Println(strings.TrimSpace(helpText))
}
