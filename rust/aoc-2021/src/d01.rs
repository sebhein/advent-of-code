use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let mut window = [0, 0, 0];
    let mut increases = -3;
    if let Ok(lines) = read_lines("../data/01") {
        for line in lines {
            if let Ok(depth) = line {
                let d = depth.parse().unwrap();

                let old_sum = window[0] + window[1] + window[2];
                window = [window[1], window[2], d];
                let new_sum = window[0] + window[1] + window[2];

                if new_sum > old_sum {
                    increases += 1
                }
            }
        }
    }

    println!("increases: {}", increases);
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
