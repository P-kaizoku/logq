# logq

Streaming web-server log analyzer in Go. Reads logs from stdin, counts fields, prints the top-N — all in a single pass.

## Features

- Counts occurrences of any field and prints the top-N entries
- `-field ip|status` selects what to count
- `-n` controls how many results
- Streams line-by-line with `bufio.Scanner` (no full-file load)
- Allocation-free IP extraction via `strings.IndexByte` + slicing
- Status parsing via `strings.FieldsSeq` (Go 1.24 lazy iterator) with early exit

## Usage

```bash
cat access.log | logq -field ip -n 10       # top 10 IPs by request count
cat access.log | logq -field status -n 5    # top 5 status codes
```

## Sample

```bash
$ logq -field status -n 5 < testdata/gen.log
200 --> 2220570
301 --> 444953
500 --> 444951
404 --> 444782
401 --> 444744
```

## Benchmark

Parses a 4M-line (362MB) generated nginx-style log file in under 2 seconds:

```
time go run . -field status -n 5 < testdata/gen.log
real 0m1.456s
```

The `strings.FieldsSeq`-based status parser is notably faster than building a full field slice per line, since it can return early on the first 3-digit token.

## Data

Generate test data with the bundled generator:

```bash
make logs          # writes 4M lines to testdata/gen.log
```

## Design notes

- The status code is located by **rule, not position**: the only 3-digit all-digit token on the line. Position-based parsing silently corrupts when log format changes — a real bug I fixed by parsing what the field *is*, not where it sits.
- CLI-first design: reads stdin, writes stdout, composes with Unix pipes (`head`, `sort`, etc.).