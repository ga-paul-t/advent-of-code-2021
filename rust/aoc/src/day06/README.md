# Day 6: Lantern Fish

Read about the puzzle [here](https://adventofcode.com/2021/day/6).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.003ms**
- Part B: **0.003ms**

```sh
$ cargo +nightly run --release --bin benchmark

🧩 Puzzle 06a: 1590327954513 [time taken: 3.93 μs]
🧩 Puzzle 06b: 1590327954513 [time taken: 3.94 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- day06::
    Finished bench [optimized] target(s) in 0.02s
     Running unittests (target/release/deps/aoc-f4fcf1000ee7c6f4)

running 4 tests
test day06::a::tests::test_solve ... ignored
test day06::b::tests::test_solve ... ignored
test day06::a::tests::bench_run ... bench:       2,867 ns/iter (+/- 477)
test day06::b::tests::bench_run ... bench:       2,700 ns/iter (+/- 560)
```
