use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::VecDeque;

fn main() {
    let mut window: VecDeque<i32> = [0, 0, 0].into();
    let mut increases = -3;
    if let Ok(lines) = read_lines("../data/01") {
        for line in lines {
            if let Ok(depth) = line {
                let d_int = depth.parse().unwrap();

                let old_sum: i32 = window.iter().sum();
                window.pop_front();
                window.push_back(d_int);
                let new_sum: i32 = window.iter().sum();

                if new_sum > old_sum {
                    increases += 1
                }
            }
        }
    }

    println!("step 2, increases: {}", increases);
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
