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

- Part A: **0.218ms**
- Part B: **0.538ms**

```sh
$ cargo +nightly run --release --bin benchmark

🎄 Advent of Code 2021 - Benchmark [executions: 200]

🧩 Puzzle 05a: 5169  [time taken: 218.97 μs]
🧩 Puzzle 05b: 22083 [time taken: 538.66 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- day05::
    Finished bench [optimized] target(s) in 0.02s
     Running unittests (target/release/deps/aoc-f4fcf1000ee7c6f4)

running 4 tests
test day05::a::tests::test_solve ... ignored
test day05::b::tests::test_solve ... ignored
test day05::a::tests::bench_run ... bench:     213,204 ns/iter (+/- 35,523)
test day05::b::tests::bench_run ... bench:     575,694 ns/iter (+/- 78,024)
```
