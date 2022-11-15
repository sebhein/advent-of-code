use std::env;
use std::process;

use aoc23::Day;

fn main() {
    let args: Vec<String> = env::args().collect();

    let day = Day::build(&args).unwrap_or_else(|err| {
        println!("Problem parsing args: {err}");
        process::exit(1);
    });

    if let Err(e) = aoc23::solve(day) {
        println!("Application error: {e}");
        process::exit(1);
    };
}

