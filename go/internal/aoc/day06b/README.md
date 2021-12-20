# Day 6: Lanternfish - Part B

Read about the puzzle [here](https://adventofcode.com/2021/day/5).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Puzzle was run 200 times and the fastest time was recorded as: **0.006ms**

```sh
$ aoc-2021 --benchmark --puzzle 6

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 06b: 1590327954513 [time taken: 6.577µs]
```

## Go Benchmark

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_Run$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day06b

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day06b
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	  221930	      5338 ns/op	    4864 B/op	       1 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day06b	1.360s
```
