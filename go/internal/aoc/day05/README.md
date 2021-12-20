# Day 5: Hydrothermal Venture

Read about the puzzle [here](https://adventofcode.com/2021/day/5).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.809ms**
- Part B: **1.232ms**

```sh
$ aoc-2021 --benchmark --puzzle 5

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 05a: 5169  [time taken: 0.809µs]
🧩 Puzzle 05b: 22083 [time taken: 1.232ms]
```

## Go Benchmark

Puzzle A:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleA$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16            1166        882285 ns/op     8076170 B/op       1503 allocs/op
PASS
ok      github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05    1.327s
```

Puzzle B:

```sh
$ go test -benchmem -run=^$ -bench ^Benchmark_PuzzleB$ github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05

goos: darwin
goarch: amd64
pkg: github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Run-16             823       1444289 ns/op     8076174 B/op       1503 allocs/op
PASS
ok      github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05    1.603s
```
