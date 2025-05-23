# matomail

`matomail` is a command-line utility that combines the concept of "matome" (Japanese for "gather/collect") with the functionality similar to the Unix "tail" command.

* **matome** (纏める): to gather, collect, summarize, or consolidate information
* **Purpose**: Efficiently condense log output by grouping consecutive identical lines, making log analysis easier

# USAGE

Basic usage involves piping output from another command (such as `tail`) into `matomail`:

```
% tail /var/log/foo.log | matomail
2025-05-14 20:48:48 FOO
(1) 2025-05-14 20:48:49 BAR
...
(3) 2025-05-14 20:48:51 BAR
%
```

Compare with the original output:

```
% tail /var/log/foo.log
2025-05-14 20:47:48 FOO
2025-05-14 20:48:49 BAR
2025-05-14 20:48:50 BAR
2025-05-14 20:48:51 BAR
```

Notice how the three consecutive "BAR" lines are condensed into a more readable format.

# SPECIFICATION

- Reads lines from STDIN in real-time
- Operates without buffering for immediate output
- When a new line is identical to the previous line:
  - Clears the current line
  - Adds a count with parenthesis prefix (e.g., "(3)")
  - Prints the line immediately after the prefix
- Displays usage information with the `--help` option
- Customizable "same as previous" comparison logic:
  - By default, lines are considered identical when they match exactly
  - If a regular expression is specified, matching portions are masked during comparison
  - Example: With regex `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, the timestamp portion of lines is masked when comparing, allowing lines with different timestamps but identical content to be grouped together
- Handling of consecutive identical lines:
  - For the first occurrence, displays the line with a `(1)` prefix
  - For subsequent occurrences, replaces intermediate lines with `...`
  - For the last occurrence, displays the line with the final count as prefix (e.g., `(42)`)

# TEST

Example of basic functionality:

```
% cat test_data/1
asdf
asdf
asdf
% cat test_data/1 | ./matomail
(1) asdf
...
(3) asdf
```

This demonstrates how three identical lines are condensed into a more readable format with a counter.

# RUNNING TESTS

To run the tests:

```
% cd src
% go test
```

For more verbose output:

```
% cd src
% go test -v
```

The test suite includes:
1. Basic functionality test - verifies handling of consecutive identical lines
2. Regex pattern test - tests masking of timestamps and other dynamic content
3. Empty input test - ensures proper handling of empty input
4. Special characters test - validates handling of lines with special characters

Note: The current tests use mock implementations. In the future, these can be updated to test against the actual binary.
