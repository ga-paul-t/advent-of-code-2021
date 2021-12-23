# Day 7: The Treachery of Whales

Read about the puzzle [here](https://adventofcode.com/2021/day/7).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.017ms**
- Part B: **0.016ms**

```sh
$ cargo +nightly run --release --bin benchmark

🧩 Puzzle 07a: 329389   [time taken: 17.49 μs]
🧩 Puzzle 07b: 86397080 [time taken: 16.95 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- day07::
    Finished bench [optimized] target(s) in 0.02s
     Running unittests (target/release/deps/aoc-f4fcf1000ee7c6f4)

running 6 tests
test day07::a::tests::test_example ... ignored
test day07::a::tests::test_puzzle ... ignored
test day07::b::tests::test_example ... ignored
test day07::b::tests::test_puzzle ... ignored
test day07::a::tests::bench_run ... bench:      19,613 ns/iter (+/- 2,770)
test day07::b::tests::bench_run ... bench:      19,527 ns/iter (+/- 5,957)
```
