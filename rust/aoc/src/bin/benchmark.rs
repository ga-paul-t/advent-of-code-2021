use std::ops::Add;
use std::time::Duration;
use took::{Timer, Took};

const RUNS: usize = 200;

fn main() {
    println!("🎄 Advent of Code 2021 - Benchmark [executions: {}]", RUNS);
    println!();

    let executions: Vec<_> = aoc::jobs()
        .iter()
        .map(|job| (
            job.1,
            job.0(),
            (0..RUNS).map(|_| {
                let now = Timer::new();
                job.0();
                now.took().into_std()
            }).min().unwrap(),
        ))
        .collect();

    executions.iter().for_each(|exc| {
        println!("🧩 Puzzle {}: {:<14} [time taken: {}]", exc.0, exc.1, Took::from_std(exc.2))
    });

    let total_elapsed = executions.iter()
        .fold(Duration::new(0,0), |total, exec | total.add(exec.2));

    println!();
    println!("Total elapsed time: [{}]", Took::from_std(total_elapsed));
}