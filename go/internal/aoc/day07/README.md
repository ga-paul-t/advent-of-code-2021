# Day 7: The Treachery of Whales

Read about the puzzle [here](https://adventofcode.com/2021/day/7).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.076ms**
- Part B: **0.019ms**

```sh
$ aoc-2021 --benchmark --puzzle 7

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 07a: 329389   [time taken: 76.414µs]
🧩 Puzzle 07b: 86397080 [time taken: 19.01µs]
```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day07

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day07
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PuzzleA-16    	   15303	     78929 ns/op	   24600 B/op	       3 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day07	2.332s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day07

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day07
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PuzzleB-16    	   53695	     22601 ns/op	   24576 B/op	       2 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day07	1.560s
```
