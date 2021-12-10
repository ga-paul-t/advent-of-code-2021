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
| Day 1 | [`0.046ms`](./go/internal/aoc/day01a/task.go) | [`0.036ms`](./go/internal/aoc/day01b/task.go) | Day 1 | [`0.041ms`](./rust/day01a/src/main.rs) | [`0.040ms`](./rust/day01b/src/main.rs) |
