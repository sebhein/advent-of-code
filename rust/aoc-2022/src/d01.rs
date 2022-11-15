use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let mut current = 0;
    // This whole thing is gross and can probably be done
    // using a vector and sorting
    let mut max_first = 0;
    let mut max_second = 0;
    let mut max_third = 0;
    if let Ok(lines) = read_lines("../data/01") {
        for line in lines {
            if let Ok(cal) = line {
                if cal.chars().count() > 1 {
                    let calorie: i32 = cal.parse().unwrap();
                    current += calorie;
                } else {
                    if current > max_first {
                        max_third = max_second;
                        max_second = max_first;
                        max_first = current;
                    } else if current > max_second {
                        max_third = max_second;
                        max_second = current;
                    } else if current > max_third {
                        max_third = current;
                    }
                    current = 0;
                }
            }
        }
    }
    if current > max_first {
        max_third = max_second;
        max_second = max_first;
        max_first = current;
    } else if current > max_second {
        max_third = max_second;
        max_second = current;
    } else if current > max_third {
        max_third = current;
    }

    println!("{}", max_first + max_second + max_third);
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
