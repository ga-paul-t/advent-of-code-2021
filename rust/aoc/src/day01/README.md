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

Both puzzles were run 100 times and the fastest time was recorded as:

- Part A: **0.041ms**
- Part B: **0.040ms**

```sh
$ cargo +nightly run --release --bin benchmark

🎄 Advent of Code 2021 - Benchmark

🧩 Puzzle 01a: 1298 [time taken: 41.76 μs]
🧩 Puzzle 01b: 1248 [time taken: 40.19 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- day01::
    Finished bench [optimized] target(s) in 0.00s
     Running unittests (target/release/deps/aoc-477aaeba5962c117)

running 4 tests
test day01::a::tests::test_solve ... ignored
test day01::b::tests::test_solve ... ignored
test day01::a::tests::bench_run ... bench:      33,969 ns/iter (+/- 2,834)
test day01::b::tests::bench_run ... bench:      34,910 ns/iter (+/- 5,159)
```
