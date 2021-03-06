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

| Go    | Part A                                    | Part B                                    | Rust  | Part A                                 | Part B                                 |
| :---- | :---------------------------------------- | :---------------------------------------- | :---- | :------------------------------------- | :------------------------------------- |
| Day 1 | [`0.046ms`](./go/internal/aoc/day01/a.go) | [`0.036ms`](./go/internal/aoc/day01/b.go) | Day 1 | [`0.034ms`](./rust/aoc/src/day01/a.rs) | [`0.035ms`](./rust/aoc/src/day01/b.rs) |
| Day 2 | [`0.089ms`](./go/internal/aoc/day02/a.go) | [`0.091ms`](./go/internal/aoc/day02/b.go) | Day 2 | [`0.041ms`](./rust/aoc/src/day02/a.rs) | [`0.041ms`](./rust/aoc/src/day02/b.rs) |
| Day 3 | [`0.026ms`](./go/internal/aoc/day03/a.go) | [`0.058ms`](./go/internal/aoc/day03/b.go) | Day 3 | [`0.028ms`](./rust/aoc/src/day03/a.rs) | [`0.032ms`](./rust/aoc/src/day03/b.rs) |
| Day 4 | [`0.248ms`](./go/internal/aoc/day04/a.go) | [`0.443ms`](./go/internal/aoc/day04/b.go) | Day 4 | -                                      | -                                      |
| Day 5 | [`0.809ms`](./go/internal/aoc/day05/a.go) | [`1.232ms`](./go/internal/aoc/day05/b.go) | Day 5 | [`0.218ms`](./rust/aoc/src/day05/a.rs) | [`0.538ms`](./rust/aoc/src/day05/b.rs) |
| Day 6 | [`0.004ms`](./go/internal/aoc/day06/a.go) | [`0.006ms`](./go/internal/aoc/day06/b.go) | Day 6 | [`0.003ms`](./rust/aoc/src/day06/a.rs) | [`0.003ms`](./rust/aoc/src/day06/b.rs) |
| Day 7 | [`0.076ms`](./go/internal/aoc/day07/a.go) | [`0.019ms`](./go/internal/aoc/day07/b.go) | Day 7 | [`0.017ms`](./rust/aoc/src/day07/a.rs) | [`0.016ms`](./rust/aoc/src/day07/b.rs) |
| Day 8 | [`0.071ms`](./go/internal/aoc/day08/a.go) | [`0.147ms`](./go/internal/aoc/day08/b.go) | Day 8 | [`0.023ms`](./rust/aoc/src/day08/a.rs) | [`0.059ms`](./rust/aoc/src/day08/b.rs) |
| Day 9 | [`0.087ms`](./go/internal/aoc/day09/a.go) | [`0.110ms`](./go/internal/aoc/day09/b.go) | Day 9 | [`0.056ms`](./rust/aoc/src/day09/a.rs) | [`0.084ms`](./rust/aoc/src/day09/b.rs) |
