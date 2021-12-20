# Day 1: Sonar Sweep

Read about the puzzle [here](https://adventofcode.com/2021/day/1).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.035ms**
- Part B: **0.036ms**

```sh
$ aoc-2021 --benchmark --puzzle 1

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 01a: 1298 [time taken: 35.851µs]
🧩 Puzzle 01b: 1248 [time taken: 36.787µs]


```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	   36109	     33255 ns/op	   32768 B/op	       1 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01	1.658s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	   33440	     35628 ns/op	   32768 B/op	       1 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01	1.734s
```
