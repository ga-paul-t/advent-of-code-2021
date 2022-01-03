# Day 9: Smoke Basin

Read about the puzzle [here](https://adventofcode.com/2021/day/9).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.087ms**
- Part B: **0.110ms**

```sh
$ aoc-2021 --benchmark --puzzle 8

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 09a: 480     [time taken: 87.343µs]
🧩 Puzzle 09b: 1045660 [time taken: 110.411µs]
```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day09

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day09
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PuzzleA-16    	   12532	     96216 ns/op	   27784 B/op	     428 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day09	2.301s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day09

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day09
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PuzzleB-16    	    9092	    132919 ns/op	   98192 B/op	     112 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day09	2.365s
```
