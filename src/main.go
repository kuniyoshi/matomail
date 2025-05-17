package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

const (
	upOneLine = "\033[1A"
	eraseLine = "\033[2K"
)

type LineProcessor struct {
	Writer  *bufio.Writer
	Scanner *bufio.Scanner
	Re      *regexp.Regexp
	Debug   bool
}

func NewLineProcessor(writer *bufio.Writer, scanner *bufio.Scanner, re *regexp.Regexp, debug bool) *LineProcessor {
	return &LineProcessor{
		Writer:  writer,
		Scanner: scanner,
		Re:      re,
		Debug:   debug,
	}
}

func (lp *LineProcessor) ProcessLines() error {
	var prevLine string
	var currentLine string
	var count int

	for lp.Scanner.Scan() {
		prevLine = currentLine
		currentLine = lp.Scanner.Text()

		isSame := areLinesSame(currentLine, prevLine, lp.Re)

		if lp.Debug {
			fmt.Fprintf(os.Stderr, "DEBUG: isSame: %v\n", isSame)
		}

		if !isSame {
			count = 0
			fmt.Fprintf(lp.Writer, "%s\n", currentLine)
			lp.Writer.Flush()
			continue
		}

		count++

		if count == 1 {
			fmt.Fprint(lp.Writer, upOneLine)
			fmt.Fprint(lp.Writer, eraseLine)
			fmt.Fprintf(lp.Writer, "(1) %s\n", prevLine)
			fmt.Fprintf(lp.Writer, "(2) %s\n", currentLine)
			lp.Writer.Flush()
		} else {
			fmt.Fprint(lp.Writer, upOneLine)
			fmt.Fprint(lp.Writer, eraseLine)
			fmt.Fprintf(lp.Writer, "...\n")
			fmt.Fprintf(lp.Writer, "(%d) %s\n", count+1, currentLine)
			lp.Writer.Flush()
		}
	}

	if err := lp.Scanner.Err(); err != nil {
		return fmt.Errorf("error reading from scanner: %w", err)
	}

	return nil
}

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

	stdout := bufio.NewWriterSize(os.Stdout, 1) // 出力をバッファリングしないように設定
	scanner := bufio.NewScanner(os.Stdin)

	processor := NewLineProcessor(stdout, scanner, re, true)
	if err := processor.ProcessLines(); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing lines: %v\n", err)
		os.Exit(1)
	}
}

func areLinesSame(currentLine, prevLine string, re *regexp.Regexp) bool {
	if re == nil {
		return currentLine == prevLine
	}
	return re.ReplaceAllString(currentLine, "") == re.ReplaceAllString(prevLine, "")
}

// 使用方法を表示する関数
func displayHelp() {
	fmt.Print(`
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

EXAMPLES:
  tail -f /var/log/syslog | matomail
  tail -f /var/log/syslog | matomail --pattern="\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}"
`)
}
