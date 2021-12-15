# Day 3: Binary Diagnostic - Part B

Read about the puzzle [here](https://adventofcode.com/2021/day/3).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Puzzle was run 200 times and the fastest time was recorded as: **0.058ms**

```sh
$ aoc-2021 --benchmark --puzzle 3

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 03b: 3385170 [time taken: 58.824µs]
```

## Go Benchmark

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_Run$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03a

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03a
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	   56538	     20684 ns/op	   16448 B/op	       5 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03a	1.511s
```
