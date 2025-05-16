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
- `matome` gathers same lines
  - The first line of the lines, it has a `(1)` prefix, and leads space
  - The last line of the lines, it has a counted prefix
  - Other lines are gathered as `...`

# TEST

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

# TODO

Based on the specification, the following features need to be implemented:

1. Core functionality:
   - [ ] Read lines from STDIN without buffering
   - [ ] Implement line comparison logic
   - [ ] Add counter for identical consecutive lines
   - [ ] Clear and rewrite lines with counter prefix

2. Command-line options:
   - [ ] Implement `--help` option to display usage information
   - [ ] Add option for specifying custom regex pattern for line comparison

3. Performance optimizations:
   - [ ] Ensure efficient handling of high-volume log streams
   - [ ] Minimize memory usage for long-running processes

4. Testing:
   - [ ] Create unit tests for line comparison logic
   - [ ] Add integration tests with sample log files
   - [ ] Test with various regex patterns for comparison customization

5. Documentation:
   - [ ] Add detailed documentation for all command-line options
   - [ ] Include examples of common use cases
   - [ ] Document performance characteristics and limitations
