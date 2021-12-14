use std::ops::Add;
use std::time::Duration;
use took::{Timer, Took};

fn main() {
    println!("🎄 Advent of Code 2021");
    println!();

    let mut total_elapsed = Duration::new(0, 0);
    aoc::jobs()
        .iter()
        .for_each(|job| {
            let now = Timer::new();
            let result = job.0();
            let taken = now.took().into_std();
            println!("🧩 Puzzle {}: {:<14} [time taken: {}]", job.1, result, Took::from_std(taken));

            total_elapsed = total_elapsed.add(taken);
        });

    println!();
    println!("Total elapsed time: [{}]", Took::from_std(total_elapsed));
}