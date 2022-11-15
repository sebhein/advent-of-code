pub fn solve(_input: &String) {
    let sol_1 = _solve_p1(_input);
    println!("day 02, part 1, {sol_1}");
    let sol_2 = _solve_p2(_input);
    println!("day 01, part 2, {sol_2}");
}

fn _solve_p1(_input: &String) -> i32 {
    let mut sum = 0;
    let max = vec![12, 13, 14];
    for line in _input.lines() {
        let game = line.split(":").collect::<Vec<&str>>();
        let id = game[0].split(" ").collect::<Vec<&str>>()[1];
        let reveals = game[1].trim().split(";").collect::<Vec<&str>>();
        // red, green, blue
        let mut max_cubes = vec![0, 0, 0];
        for reveal in reveals {
            let cubes = reveal.split(",").collect::<Vec<&str>>();
            for cube in cubes {
                let count_and_color = cube.trim().split(" ").collect::<Vec<&str>>();
                let count = count_and_color[0].parse::<i32>().unwrap();
                let color = count_and_color[1];
                let index = match color {
                    "red" => 0,
                    "green" => 1,
                    "blue" => 2,
                    _ => 3,
                };
                if max_cubes[index] < count {
                    max_cubes[index] = count;
                }
            }
        }
        
        let mut possible = true;
        for i in 0..3 {
            if max_cubes[i] > max[i] {
                possible = false;
            }
        }
        if possible {
            sum += id.parse::<i32>().unwrap();
        }
    }
    return sum;
}


fn _solve_p2(_input: &String) -> i32 {
    let mut sum = 0;
    for line in _input.lines() {
        let game = line.split(":").collect::<Vec<&str>>();
        let reveals = game[1].trim().split(";").collect::<Vec<&str>>();
        // red, green, blue
        let mut max_cubes = vec![0, 0, 0];
        for reveal in reveals {
            let cubes = reveal.split(",").collect::<Vec<&str>>();
            for cube in cubes {
                let count_and_color = cube.trim().split(" ").collect::<Vec<&str>>();
                let count = count_and_color[0].parse::<i32>().unwrap();
                let color = count_and_color[1];
                let index = match color {
                    "red" => 0,
                    "green" => 1,
                    "blue" => 2,
                    _ => 3,
                };
                if max_cubes[index] < count {
                    max_cubes[index] = count;
                }
            }
        }
        
        sum += max_cubes.iter().fold(1, |acc, x| acc * x);
    }
    return sum;
}
