# Day 2: Dive!

Read about the puzzle [here](https://adventofcode.com/2021/day/2).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.088ms**
- Part B: **0.088ms**

```sh
$ aoc-2021 --benchmark --puzzle 2

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 02a: 1990000    [time taken: 88.213µs]
🧩 Puzzle 02b: 1975421260 [time taken: 88.776µs]
```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	   15096	     79688 ns/op	   48384 B/op	    1001 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02	2.154s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16           15066         81219 ns/op       48384 B/op       1001 allocs/op
PASS
ok      github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02    2.152s
```
