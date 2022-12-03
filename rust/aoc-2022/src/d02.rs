use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::collections::HashMap;


// A = Rock
// B = Paper
// C = Scissor
//
// Part 1
// X = Rock
// Y = Paper
// Z = Scissor
//
// Part 2
// X = lose
// Y = draw
// Z = win


fn main() {
    // feels like cheating
    let combos: HashMap<&str, i32> = HashMap::from([
        ("A X", 3),  // loss / scissor -> 0 + 3
        ("A Y", 4),  // draw / rock -> 3 + 1
        ("A Z", 8),  // win / paper -> 6 + 2

        ("B X", 1),  // loss / rock -> 0 + 1
        ("B Y", 5),  // draw / paper -> 3 + 2
        ("B Z", 9),  // win / scissor -> 6 + 3

        ("C X", 2),  // loss / paper -> 0 + 2
        ("C Y", 6),  // draw / scissor -> 3 + 3
        ("C Z", 7),  // win / rock -> 6 + 1
    ]);
    let mut score = 0;
    if let Ok(lines) = read_lines("../data/02") {
        for line in lines {
            if let Ok(strat) = line {
                //let round = strat.split_whitespace().collect::<Vec<&str>>();
                //match round[1] {
                    //"X" => score += 1,
                    //"Y" => score += 2,
                    //"Z" => score += 3,
                    //&_ => panic!("player move not recognized"),
                //}
                //score += combos[&*round.join("")];
                score += combos[&*strat];
            }
        }
    }

    println!("{}", score);
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
