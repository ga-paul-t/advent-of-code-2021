# Advent of Code 2021

Here are my solutions for the [Advent of Code 2021](https://adventofcode.com/2021) Challenge. Solutions are to be provided in Go and Rust<sup>1</sup>.

I will attempt to write as elegant and performant a solution as possible<sup>2</sup> in both languages. Is it possible to solve every puzzle in under a second? All language features are on the table.

<sup>1. I have never used Rust at the time of writing, but feel this is a perfect opportunity to learn</sup>
</br>
<sup>2. I don't profess to being an expert in any of these languages</sup>

## Quick Start

To build and run Go solutions:

```sh
cd go && make
./bin/aoc-2021 --benchmark
```

To build and run Rust solutions:

```sh
cd rust/aoc
cargo +nightly run --bin benchmark --release
```

## Results

All puzzles were benchmarked on the following laptop:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

| Go    | Part A                                        | Part B                                        | Rust  | Part A                                 | Part B                                 |
| :---- | :-------------------------------------------- | :-------------------------------------------- | :---- | :------------------------------------- | :------------------------------------- |
| Day 1 | [`0.046ms`](./go/internal/aoc/day01a/task.go) | [`0.036ms`](./go/internal/aoc/day01b/task.go) | Day 1 | [`0.034ms`](./rust/aoc/src/day01/a.rs) | [`0.035ms`](./rust/aoc/src/day01/b.rs) |
| Day 2 | [`0.089ms`](./go/internal/aoc/day02a/task.go) | [`0.091ms`](./go/internal/aoc/day02b/task.go) | Day 2 | [`0.041ms`](./rust/aoc/src/day02/a.rs) | [`0.041ms`](./rust/aoc/src/day02/b.rs) |
| Day 3 | [`0.026ms`](./go/internal/aoc/day03a/task.go) | [`0.058ms`](./go/internal/aoc/day03b/task.go) | Day 3 | [`0.028ms`](./rust/aoc/src/day03/a.rs) | [`0.032ms`](./rust/aoc/src/day03/b.rs) |
| Day 4 | [`0.248ms`](./go/internal/aoc/day04a/task.go) | [`0.443ms`](./go/internal/aoc/day04b/task.go) | Day 4 | -                                      | -                                      |
| Day 5 | [`0.809ms`](./go/internal/aoc/day05a/task.go) | [`1.232ms`](./go/internal/aoc/day05b/task.go) | Day 5 | [`0.218ms`](./rust/aoc/src/day05/a.rs) | [`0.538ms`](./rust/aoc/src/day05/b.rs) |
| Day 6 | [`0.022ms`](./go/internal/aoc/day06a/task.go) | [`0.046ms`](./go/internal/aoc/day06b/task.go) | Day 6 | [`0.003ms`](./rust/aoc/src/day06/a.rs) | [`0.003ms`](./rust/aoc/src/day06/b.rs) |
