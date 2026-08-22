# Level 01 - Go foundations deep (Weeks 1-2)

You "know Go." This level removes the places where that is a lie. Interviewers probe exactly these seams with juniors.

## Objectives

1. Explain and predict memory/aliasing behavior of slices and maps.
2. Use interfaces without falling into the nil-interface trap.
3. Handle errors like the stdlib does: wrap, inspect, decide.
4. Operate tooling fluently: modules, vet, fmt, race detector, benchmarks.

## Concepts

- Slices: len vs cap, growth factor, `append` aliasing bugs, slicing a slice shares backing array, full slice expression `s[low:high:max]`, `copy` semantics
- Maps: no addressable values, iteration randomness, zero-value gotchas, when to pre-size
- Interfaces: implicit satisfaction, nil interface vs typed-nil pointer, type assertions and `switch`, small-interface design (io.Reader/Writer)
- Errors: `fmt.Errorf` + `%w`, `errors.Is/As`, sentinel errors, typed errors, when to panic (almost never)
- Structs: value vs pointer receivers and method sets, embedding vs inheritance, struct tags
- Generics: constraints, when generics beat `interface{}`/any
- Packages: internal/, export rules, init() ordering
- Tooling: go.mod hygiene, `go vet`, `gofmt`/`goimports`, `-race`, `go test -bench`

## Build tasks

1. From a blank file, no tutorial: CLI log analyzer (`logq`) - parse nginx-ish logs from stdin, output top-N IPs/status codes, streaming with bufio. Uses maps, slices, sorting, flags, error wrapping.
2. Write deliberately broken programs demonstrating each trap below; fix them; keep in `levels/01-go-foundations/traps/`.

## Exercises

- Predict-then-run: 10 snippet drills on slice aliasing + append capacity. I supply snippets live.
- Implement `func Dedupe[T comparable](in []T) []T` two ways (map-based, in-place for sorted).
- Benchmark map presizing vs not; write one paragraph on results in README.

## Checkpoint (demonstrate to advance)

- [ ] Whiteboard-explain: what `append` does when cap is exceeded; why iterating a map is random; the typed-nil-in-interface trap.
- [ ] logq works on a 100MB generated log file under ~2s; benchmark table committed.
- [ ] All trap programs fixed with written one-line explanations.
- [ ] DSA: >=12 new NeetCode problems solved in Go this level.

## Resources

- The Go Tour (skim only where weak): go.dev/tour
- Effective Go: go.dev/doc/effective_go
- Go slices usage and internals: go.dev/blog/slices-intro and go.dev/src/runtime/slice.go
- Book: 100 Go Mistakes and How to Avoid Them - chapters 1-8 (this book serves all levels)
