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

- Part A: **0.034ms**
- Part B: **0.035ms**

```sh
$ cargo +nightly run --release --bin benchmark

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 01a: 1298 [time taken: 34.88 μs]
🧩 Puzzle 01b: 1248 [time taken: 35.26 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- day02::
    Finished bench [optimized] target(s) in 0.00s
     Running unittests (target/release/deps/aoc-477aaeba5962c117)

running 4 tests
test day02::a::tests::test_solve ... ignored
test day02::b::tests::test_solve ... ignored
test day02::a::tests::bench_run ... bench:      38,126 ns/iter (+/- 3,896)
test day02::b::tests::bench_run ... bench:      37,579 ns/iter (+/- 3,423)
```
