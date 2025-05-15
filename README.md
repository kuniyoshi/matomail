# matomail

`matomail` combines "matome" (Japanese for "gather/collect") with "tail".

* matome (纏める): to gather, collect, or summarize

# USAGE

```
% tail /var/log/foo.log | matomail
2025-05-14 20:48:48 FOO
(1) 2025-05-14 20:48:49 BAR
...
(3) 2025-05-14 20:48:51 BAR
%
% tail /var/log/foo.log
2025-05-14 20:47:48 FOO
2025-05-14 20:48:49 BAR
2025-05-14 20:48:50 BAR
2025-05-14 20:48:51 BAR
```

# SPECIFICATION

- Reads lines from STDIN
- Operates without buffering
- When a new line is identical to the previous line:
  - Clears the line
  - Adds a count with parenthesis prefix (e.g., "(3)")
  - Prints the line immediately after the prefix
- Displays usage information with the `--help` option
- Customizable "same as previous" comparison logic:
  - By default, lines are considered identical when they match exactly
  - If a regular expression is specified, matching portions are masked during comparison
  - Example: With regex `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, the datetime portion of lines is masked when comparing
