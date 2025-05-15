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
	// コマンドラインオプションの定義
	helpFlag := flag.Bool("help", false, "Display usage information")
	regexPattern := flag.String("pattern", "", "Regular expression pattern for line comparison masking")
	flag.Parse()

	// --helpオプションが指定された場合、使用方法を表示して終了
	if *helpFlag {
		displayHelp()
		return
	}

	// 正規表現パターンが指定された場合、コンパイル
	var re *regexp.Regexp
	if *regexPattern != "" {
		var err error
		re, err = regexp.Compile(*regexPattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid regular expression pattern: %v\n", err)
			os.Exit(1)
		}
	}

	// 出力をバッファリングしないように設定
	stdout := bufio.NewWriterSize(os.Stdout, 1)

	// 標準入力からの読み取り
	scanner := bufio.NewScanner(os.Stdin)

	var prevLine string         // 前の行
	var maskedPrevLine string   // マスクされた前の行
	var count int               // 同一行のカウンター
	var isFirstLine bool = true // 最初の行かどうか

	// 標準入力から行を読み取り
	for scanner.Scan() {
		currentLine := scanner.Text()

		// 正規表現パターンが指定されている場合、マスクを適用
		maskedCurrentLine := currentLine
		if re != nil {
			maskedCurrentLine = re.ReplaceAllString(currentLine, "MASKED")
		}

		// 前の行と比較
		if maskedCurrentLine == maskedPrevLine && !isFirstLine {
			// 同一行の場合
			count++
		} else {
			// 異なる行の場合
			// 前の行があれば出力
			if !isFirstLine {
				if count > 0 {
					// カウンター付きで出力
					fmt.Fprintf(stdout, "(%d) %s\n", count, prevLine)
				} else {
					// そのまま出力
					fmt.Fprintf(stdout, "%s\n", prevLine)
				}
				stdout.Flush()
			}

			// 現在の行を保存
			prevLine = currentLine
			maskedPrevLine = maskedCurrentLine

			// カウンターをリセット
			count = 1
			isFirstLine = false
		}
	}

	// 最後の行を出力
	if !isFirstLine {
		if count > 0 {
			fmt.Fprintf(stdout, "(%d) %s\n", count, prevLine)
		} else {
			fmt.Fprintf(stdout, "%s\n", prevLine)
		}
		stdout.Flush()
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
