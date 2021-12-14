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

Both puzzles were run 100 times and the fastest time was recorded as:

- Part A: **0.033ms**
- Part B: **0.037ms**

```sh
$ cargo +nightly run --release --bin benchmark

🎄 Advent of Code 2021 - Benchmark

🧩 Puzzle 03a: 3882564 [time taken: 33.72 μs]
🧩 Puzzle 03b: 3385170 [time taken: 37.30 μs]
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
