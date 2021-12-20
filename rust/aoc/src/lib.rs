#![feature(test)]
#![feature(drain_filter)]

pub mod day01;
pub mod day02;
pub mod day03;
pub mod day04;
pub mod day05;
pub mod day06;

pub fn jobs() -> &'static [(fn() -> usize, &'static str)] {
    &[
        (day01::a::run, "01a"),
        (day01::b::run, "01b"),
        (day02::a::run, "02a"),
        (day02::b::run, "02b"),
        (day03::a::run, "03a"),
        (day03::b::run, "03b"),
        (day04::a::run, "04a"),
        (day04::b::run, "04b"),
        (day05::a::run, "05a"),
        (day05::b::run, "05b"),
        (day06::a::run, "06a"),
        (day06::b::run, "06b")
    ]
}