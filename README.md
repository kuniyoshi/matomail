# matomail

`matomail` は 「纏める」を "tail" します。

* 纏める: gather, collect

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

- Read line from STDIN
- Required to disable buffer
- If the new line is same as previous line;
  - Clear the line
  - Add count with parenthesis prefix e.g.: (3)
  - Print the line just after the prefix
- Show usage on `--help` option
- `same as previous` test is customizable
  - As a default, line will be treated as same when completely same as previous
  - If specified regular expression, line masked by it
  - e.g.: the regular expression is `\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`
          the line's datetime is masked when comparison.

