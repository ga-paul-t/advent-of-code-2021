# Day 3: Binary Diagnostic

Read about the puzzle [here](https://adventofcode.com/2021/day/3).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.026ms**
- Part B: **0.058ms**

```sh
$ aoc-2021 --benchmark --puzzle 3

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 03b: 3882564 [time taken: 26.924µs]
🧩 Puzzle 03b: 3385170 [time taken: 58.824µs]
```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	   56538	     20684 ns/op	   16448 B/op	       5 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03	1.511s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	   56538	     20684 ns/op	   16448 B/op	       5 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03	1.511s
```
