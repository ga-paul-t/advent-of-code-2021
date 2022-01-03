# Day 9: Smoke Basin

Read about the puzzle [here](https://adventofcode.com/2021/day/9).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Both puzzles were run 200 times and the fastest time was recorded as:

- Part A: **0.056ms**
- Part B: **0.084ms**

```sh
$ cargo +nightly run --release --bin benchmark

🧩 Puzzle 09a: 480     [time taken: 56.70 μs]
🧩 Puzzle 09b: 1045660 [time taken: 84.69 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- day09::
    Finished bench [optimized] target(s) in 0.02s
     Running unittests (target/release/deps/aoc-f4fcf1000ee7c6f4)

running 6 tests
test day09::a::tests::test_example ... ignored
test day09::a::tests::test_puzzle ... ignored
test day09::b::tests::test_example ... ignored
test day09::b::tests::test_puzzle ... ignored

test day09::a::tests::bench_run ... bench:      59,060 ns/iter (+/- 11,316)
test day09::b::tests::bench_run ... bench:      94,102 ns/iter (+/- 9,274)
```
