# Day 3: Binary Diagnostic

Read about the puzzle [here](https://adventofcode.com/2021/day/3).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.028ms**
- Part B: **0.032ms**

```sh
$ cargo +nightly run --release --bin benchmark

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 03a: 3882564 [time taken: 28.13 μs]
🧩 Puzzle 03b: 3385170 [time taken: 32.13 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- day03::
    Finished bench [optimized] target(s) in 0.01s
     Running unittests (target/release/deps/aoc-477aaeba5962c117)

running 4 tests
test day03::a::tests::test_solve ... ignored
test day03::b::tests::test_solve ... ignored
test day03::a::tests::bench_run ... bench:      27,922 ns/iter (+/- 3,852)
test day03::b::tests::bench_run ... bench:      32,380 ns/iter (+/- 7,951)
```
