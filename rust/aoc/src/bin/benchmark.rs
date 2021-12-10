use took::{Timer, Took};

const RUNS: usize = 100;

fn main() {
    println!("🎄 Advent of Code 2021 - Benchmark");
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
        println!("🧩 Puzzle {}: {} [time taken: {}]", exc.0, exc.1, Took::from_std(exc.2))
    })
}