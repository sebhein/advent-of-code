use std::env;
use std::process;

use aoc_2022::Day;

fn main() {
    let args: Vec<String> = env::args().collect();

    let day = Day::build(&args).unwrap_or_else(|err| {
        println!("Problem parsing args: {err}");
        process::exit(1);
    });

    if let Err(e) = aoc_2022::solve(day) {
        println!("Application error: {e}");
        process::exit(1);
    };
}

