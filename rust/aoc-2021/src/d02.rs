use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    struct Course {
        horizontal: i32,
        depth: i32,
        aim: i32,
    }

    let mut nav = Course {horizontal: 0, depth: 0, aim: 0};
    if let Ok(lines) = read_lines("../data/02") {
        for line in lines {
            if let Ok(course) = line {
                let curr = course.split_whitespace().collect::<Vec<&str>>();
                let direction = curr[0];
                let magnitude: i32 = curr[1].parse().unwrap();

                match direction {
                    "forward" => {
                        nav.horizontal += magnitude;
                        nav.depth += magnitude * nav.aim;
                    }
                    "down" => nav.aim += magnitude,
                    "up" => nav.aim -= magnitude,
                    &_ => panic!("direction not recognized"),
                }
            }
        }
    }
    println!("horizontal: {}, vertical: {}, product: {}", nav.horizontal, nav.depth, nav.horizontal * nav.depth)
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
