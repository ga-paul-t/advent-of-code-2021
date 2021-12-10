# Day 1: Sonar Sweep - Part B

Read about the puzzle [here](https://adventofcode.com/2021/day/1).

Puzzle was run on:

```text
macOS Big Sur (11.5.2)
MacBook Pro (2019)
2.3 GHz 8-Core Intel Core i9
16 GB 2667 MHz DDR4
```

## Performance

Puzzle was run 100 times and the fastest time was recorded as: **0.049ms**

```sh
$ cargo +nightly run --release --bin runner

🎄 Advent of Code 2021

🧩 Puzzle 01b: 1248 [time taken: 49.67 μs]
```

## Rust Benchmark

```sh
$ cargo bench -- tests::bench_main
    Finished bench [optimized] target(s) in 0.03s
     Running unittests (target/release/deps/day01b-f7651158934a638f)

running 1 test
test tests::bench_main ... bench:      36,258 ns/iter (+/- 8,434)

test result: ok. 0 passed; 0 failed; 0 ignored; 1 measured; 1 filtered out; finished in 0.27s
```
