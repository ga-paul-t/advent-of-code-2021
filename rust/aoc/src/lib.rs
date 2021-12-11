#![feature(test)]

pub mod day01;
pub mod day02;

pub fn jobs() -> &'static [(fn() -> usize, &'static str)] {
    &[
        (day01::a::run, "01a"),
        (day01::b::run, "01b"),
        (day02::a::run, "02a"),
        (day02::b::run, "02b")
    ]
}