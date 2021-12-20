# Day 4: Giant Squid

Read about the puzzle [here](https://adventofcode.com/2021/day/4).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.248ms**
- Part B: **0.443ms**

```sh
$ aoc-2021 --benchmark --puzzle 4

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 04a: 82440 [time taken: 248.208µs]
🧩 Puzzle 04b: 20774 [time taken: 443.738µs]
```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16            3154        351938 ns/op      235249 B/op       1289 allocs/op
PASS
ok      github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04    1.272s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16    	    1893	    601845 ns/op	  235273 B/op	    1289 allocs/op
PASS
ok  	github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04	1.325s
```
