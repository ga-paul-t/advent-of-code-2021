# Day 8: Seven Segment Search

Read about the puzzle [here](https://adventofcode.com/2021/day/8).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.071ms**
- Part B: **0.147ms**

```sh
$ aoc-2021 --benchmark --puzzle 8

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 08a: 543    [time taken: 71.823µs]
🧩 Puzzle 08b: 994266 [time taken: 147.948µs]
```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day08

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day08
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PuzzleA-16    	   15355	     78124 ns/op	   35968 B/op	     402 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day08	2.112s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day08

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day08
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PuzzleB-16    	    7410	    161542 ns/op	   54400 B/op	     601 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day08	2.342s
```
