use std::fs;
use std::error::Error;
mod d06;
mod d07;


fn not_a_day(_input: &String) {
    println!("not a valid day");
}

pub fn solve(day: Day) -> Result<(), Box<dyn Error>> {
    let contents = fs::read_to_string(day.file_path)?;

    (day.solve)(&contents);

    Ok(())
}

pub struct Day {
    solve: fn(&String),
    file_path: String,
}

impl Day{
    pub fn build(args: &[String]) -> Result<Day, &'static str> {
        if args.len() < 2 {
            return Err("woah too little args");
        }
        let day = &args[1];
        let mut file_path: String = "data/".to_owned();
        file_path.push_str(day);

        let solve = match (*day).as_str() {
            "06" => d06::solve_06,
            "07" => d07::solve_07,
            _ => not_a_day,
        };

        Ok(Day { solve, file_path })
    }
}
